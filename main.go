/*
***********************************************
Licensed under BSD-3-Clause - see license file.
***********************************************
*/

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/resty.v1"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"

	validator "github.com/xeipuuv/gojsonschema"
)

const resourcename = "resource_junos_qfx_native_bgp"
const versionminor = 12
const versionmajor = 20
const debugPrint = false

// Config used to decode configuration file
type Config struct {
	TFStaging   string `toml:"tfstaging"`
	TFProviders string `toml:"tfproviders"`
	TFUpload    string `toml:"tfupload"`
	TFDir       string `toml:"tfdir"`

	FlaskGenson string `toml:"flaskgenson"`
}

func getconfig(name string) (C Config, err error) {
	c := Config{}
	_, err = toml.DecodeFile(name, &c)

	if err != nil {
		return c, err
	}

	return c, nil
}

// GetUUID  returns a UUID
func GetUUID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Print(err)
	}
	return uuid.String()
}

// TFProcess deals with a single task pipeline
func TFProcess(cfg Config, j string, provVer string) (JSONWrapper, error) {
	// 0. jobs is the input channel for the provider name
	// 1. Copy each provider in /uploads to /tmp/XXXX where XXX is a random directory name and also /providers
	// 2. In each directory, create an empty resource and do a TF init
	// 3. Grab the schema output and whack it in ETCd
	// 4. Create a canonical example from the schema and whack that in ETCd
	// 5. Create PostGres entries for the location of the provider
	// 6. Exit the directory and rm it
	// rand.Seed(int64(id))
	// for j := range jobs {
	// 	results <- jobInfo{Worker: id, Job: j, Status: "start"}
	//
	// Create /tmp/xxxx directory
	// sleep is required to get a diff UUID
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	tmpName := GetUUID()
	tmpPath := filepath.Join(cfg.TFStaging, tmpName)

	os.MkdirAll(tmpPath, os.ModePerm)
	os.Chdir(tmpPath)

	// 2. Copy provider in /uploads to path
	sourceFile, err := os.Open(cfg.TFUpload + "/" + j)
	if err != nil {
		fmt.Println(err)
	}
	defer sourceFile.Close()
	pN := strings.Split(j, "terraform-provider-")
	provName := pN[1]
	fullProvNameSlice := strings.Split(j, "/")
	fullProvName := fullProvNameSlice[len(fullProvNameSlice)-1]
	destFile, err := os.Create(tmpPath + "/" + fullProvName)

	destFile.Chmod(0777)
	if err != nil {
		fmt.Println("error 1 ->")
		fmt.Println(err)
	}
	defer destFile.Close()
	// count, err := io.Copy(destFile, sourceFile)
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Copied %d bytes.\n", count)

	err = destFile.Sync()
	if err != nil {
		fmt.Println(err)
	}

	// 3
	// 16/09/2020: Hashicorp made some changes to the way TF plugins/providers are used. G'damn it.
	// We need to create a new file (g'damn it). I'm cheating here with the file, to prevent further parsing of the filename
	// What should happen, is we take the name of a provider "terraform-provider-junos-qfx_v13.0.0" and grab the version number.
	// From that we can populate a 'versions.tf' file in the directory with the content below:

	const versionFileContent = `terraform {
			required_providers {
			  %s = {
				source = "terrascout/providers/%s"
				version = "~> %s"
			  }
			}
			required_version = ">= 0.13.0"
		  }`

	versionsFileName := tmpPath + "/versions.tf"
	err = ioutil.WriteFile(versionsFileName, []byte(fmt.Sprintf(versionFileContent, provName, provName, provVer)), 0644)
	if err != nil {
		fmt.Print(err)
		return JSONWrapper{}, err
	}

	// 3.2 Create a directory at the following:
	// mkdir ~/.terraform.d/plugins/dave.dev/providers/%s/%s/linux_amd64
	tfdirNameTemplate := "%s/.terraform.d/plugins/terrascout/providers/%s/%s/darwin_amd64"
	tfdirName := fmt.Sprintf(tfdirNameTemplate, cfg.TFDir, provName, provVer)
	os.MkdirAll(tfdirName, os.ModePerm)

	// 3.3 Copy the provider there
	sourceFile, err = os.Open(j)
	if err != nil {
		fmt.Println(err)
	}

	// destFile2, err := os.Create(tfdirName + "/" + "terraform-provider-" + provName + "_v" + provVer) // this one breaks naming
	destFile2, err := os.Create(tfdirName + "/" + "terraform-provider-" + provName)
	if err != nil {
		fmt.Print(err)
		return JSONWrapper{}, err
	}

	destFile2.Chmod(0777)
	if err != nil {
		fmt.Println(err)
	}

	// count, err := io.Copy(destFile2, sourceFile)
	_, err = io.Copy(destFile2, sourceFile)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Copied %d bytes.\n", count)

	err = destFile2.Sync()
	if err != nil {
		fmt.Println(err)
	}

	sourceFile.Close()
	destFile2.Close()

	// 4.
	os.Chdir(tmpPath)
	cmd0 := exec.Command("terraform", "init")
	cmd0.Run()

	cmd1 := exec.Command("terraform", "providers", "schema", "-json")
	tfschemaJSON, err := cmd1.CombinedOutput()
	if err != nil {
		fmt.Println("Error creating combined output: ", err)
		return JSONWrapper{}, err
	}

	cmd1.Stderr = os.Stderr
	cmd1.Run()

	// 5. Create JSON Canonical from TF Schema
	canonicalResources, err := CreateCanonical(tfschemaJSON, provName, cfg.FlaskGenson)
	if err != nil {
		fmt.Println("Error creating Canonical Example: ", err)
		return JSONWrapper{}, err
	}

	// 7. Remove the tmp directory
	err = os.RemoveAll(tmpPath)
	if err != nil {
		fmt.Print(err)
		return JSONWrapper{}, err
	}

	return canonicalResources, nil
}

// CreateJSONSchema returns a string of JSON schema
func CreateJSONSchema(uri string, jsonStr string) (string, error) {
	// Call the Python web service (FlaskGenson) for JSON schema generation and return the result

	// Create a Resty Client
	client := resty.New()

	// just mentioning about POST as an example with simple flow
	// User Login
	resp, err := client.R().
		SetFormData(map[string]string{
			"payload": jsonStr,
		}).
		Post(uri)

	if err != nil {
		return "", err
	}

	// fmt.Println("BODY: ", string(resp.Body()))

	return string(resp.Body()), nil

}

// ValidateJSON does a validation of a JSON string against a schema
func ValidateJSON(schemadoc string, doc string) (bool, error) {
	schemaLoader := validator.NewStringLoader(schemadoc)
	docLoader := validator.NewStringLoader(doc)

	result, err := validator.Validate(schemaLoader, docLoader)
	if err != nil {
		return false, err
	}

	if !result.Valid() {
		return false, nil
	}
	return result.Valid(), err
}

// GetDetails just returns the stringified details for the module and versioning
func GetDetails() string {
	return fmt.Sprintf("{ResourceName: %s,\nJunosMinor: %d,\nJunosMajor: %d}", resourcename, versionminor, versionmajor)
}

// PrintWrapper is a dirty function that just prints to screen what it receives
func PrintWrapper(a ...interface{}) {

	if debugPrint {
		fmt.Println(a...)
	}
}

type Output struct {
	Schema     json.RawMessage
	Resources  json.RawMessage
	JSONSchema json.RawMessage
}

func NewOutput(tfschema string, data []Resource) Output {
	output := Output{}
	output.Schema = json.RawMessage(tfschema)

	// Process the resources to []byte
	resourcesStr := "["
	schemaStr := "["
	lenResources := len(data)
	resPrintCtr := 0

	for _, v := range data {

		resourcesStr += v.Instance
		schemaStr += v.JSONSchema
		if resPrintCtr < (lenResources - 1) {
			resourcesStr += ", "
			schemaStr += ", "
		}

		resPrintCtr++
	}
	resourcesStr += "]"
	schemaStr += "]"

	output.Resources = json.RawMessage(resourcesStr)
	output.JSONSchema = json.RawMessage(schemaStr)

	return output
}

func main() {

	config := flag.String("config", "", "Path to config file (required)")

	flag.Parse()

	Cfg, err := getconfig(*config)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	data, err := TFProcess(Cfg, "terraform-provider-junos-qfx", "20.32.0101")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	output := NewOutput(string(data.TFSchema), data.Resources)
	jsonData, err := json.Marshal(output)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var out bytes.Buffer
	err = json.Indent(&out, jsonData, "", "  ")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println(out.String())
}
