package main

import (
	"encoding/json"
	"errors"
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"
)

var (
	ctr int
)

const (
	attributeStrPattern    = `%q: %q`
	attributeNonStrPattern = `%q: %v`
	cfgGroupName           = "config-group-name"
)

// JSON is a simple func to test that the JSON is valid
func JSON(str string) bool {
	var jsonStr map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonStr)
	return err == nil

}

type JSONWrapper struct {
	ProviderName string
	TFSchema     []byte // This is the schema for TF
	Resources    []Resource
}

type Resource struct {
	Instance       string `json:",omitempty"`
	ResourceName   string
	ResourceSchema string // This is the generated schema from the Resource
}

// CreateCanonical receives the TFJsonSchema and spits out a canonical example
func CreateCanonical(TFjsonSchema []byte, providerName string) (JSONWrapper, error) {

	parsed := &tfjson.ProviderSchemas{}

	err := json.Unmarshal(TFjsonSchema, &parsed)
	if err != nil {
		return JSONWrapper{}, err
	}

	// out, err := json.MarshalIndent(parsed, "", "  ")
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// out = append(out, byte('\n'))

	// Build the canonical JSON plan from the schema and spit it out
	// We're going to have some fun doing this, mainly arbitrarily iterating
	// our way through the KV values for resource blocks and nested blocks.

	// This is a dirty hack because the TF marshallers are too rigid to work with.
	// The issue is, we don't know what any given schema looks like ahead of time.
	// This is dynamic in nature and thus, we need to generate canonical JSON from any schema input.
	// resourceStrings := []string{}
	resources := JSONWrapper{}
	resources.Resources = make([]Resource, 0, 1)
	depth := 0

	resources.TFSchema = TFjsonSchema
	resources.ProviderName = providerName

	for _, v1 := range parsed.Schemas {
		// First level goes through each provider_schemas

		for k2, v2 := range v1.ResourceSchemas {
			res := Resource{}
			res.ResourceName = k2

			attributeStr := `{"resource":{`
			attributeStr += `"` + k2 + `":`
			attributeStr += `{"config-group-name":{`
			// fmt.Printf("resource_schema: %v\n", k2)

			// this is the rest of our JSON string entry
			attLength := 0
			nestedLength := 0

			// Pre-check for ID
			if _, ok := v2.Block.Attributes["id"]; ok {
				attLength++
			}

			for k3, v3 := range v2.Block.Attributes {

				if k3 == "id" {
					ctr++
					goto RELOOP
				}

				// for i := 0; i < depth; i++ {
				// 	fmt.Print("\t")
				// }

				// Attributes for top layer block
				switch v3.AttributeType.FriendlyName() {
				case "string":
					if v3.Computed {
						attributeStr += fmt.Sprintf(attributeStrPattern, k3, "computed")
					} else if k3 == "resource_name" {
						attributeStr += fmt.Sprintf(attributeStrPattern, k3, cfgGroupName)
					} else {
						attributeStr += fmt.Sprintf(attributeStrPattern, k3, fmt.Sprintf("foo%d", ctr))
					}
				case "bool":
					attributeStr += fmt.Sprintf(attributeNonStrPattern, k3, "false")
				case "number":
					attributeStr += fmt.Sprintf(attributeNonStrPattern, k3, ctr)
				case "list of string":
					attributeStr += "\"" + k3 + "\"" + ": ["
					attributeStr += fmt.Sprintf("\"bar%d\", \"barbar%d\"", ctr, ctr)
					attributeStr += "]"
				case "map of string":
					attributeStr += "\"" + k3 + "\"" + ": {"
					attributeStr += fmt.Sprintf("\"foo%d\": \"bar%d\", \"foofoo%d\": \"barbar%d\"", ctr, ctr, ctr, ctr)
					attributeStr += "}"
				case "list of number":
					attributeStr += fmt.Sprintf(attributeNonStrPattern, k3, ctr)
				case "map of number":
					attributeStr += fmt.Sprintf(attributeNonStrPattern, k3, ctr)
				case "list of bool":
					attributeStr += "["
					attributeStr += fmt.Sprintf(attributeNonStrPattern, k3, "true, false, true")
					attributeStr += "]"
				case "map of bool":
					attributeStr += "{"
					attributeStr += fmt.Sprintf(attributeNonStrPattern, k3, "\"thing%d\": true, \"thingy%d\": false")
					attributeStr += "}"
				default:
					panic(fmt.Sprintf("caught exception: %v, %v ", k3, v3))
				}

				attLength++

				if attLength < len(v2.Block.Attributes) {
					attributeStr += ", "
				} else {
					if len(v2.Block.NestedBlocks) != 0 {
						attributeStr += ", "
					}
				}

				ctr++

				//	This is tactically placed. Helps us to debug for issues
			RELOOP:
				// fmt.Printf("attributes: %v, schema: %v\n", k3, v3.AttributeType.FriendlyName())
			}

			for k4, v4 := range v2.Block.NestedBlocks {
				// for i := 0; i < depth; i++ {
				// 	fmt.Print("\t")
				// }

				// fmt.Printf("nested block: %v, type: %v\n", k4, v4.NestingMode)

				switch v4.NestingMode {
				case "single", "map", "set":
					attributeStr += fmt.Sprintf("%q: {", k4)
				case "list":
					attributeStr += fmt.Sprintf("%q: [{", k4)
				}

				printNested(v4, depth, &attributeStr)

				switch v4.NestingMode {
				case "single", "map", "set":
					attributeStr += "}"
				case "list":
					attributeStr += "}]"
				}

				nestedLength++
				if nestedLength < len(v2.Block.NestedBlocks) {
					attributeStr += ", "
				}

			}
			attributeStr += "}}}}"
			// resourceStrings = append(resourceStrings, attributeStr)
			res.Instance = attributeStr

			resources.Resources = append(resources.Resources, res)
		}
	}

	// Exit point for processing
	// fmt.Println()

	// overall := true

	// Let's do some basic JSON validation
	for _, v := range resources.Resources {
		result := JSON(v.Instance)
		// if result {
		// 	fmt.Println("Valid JSON: ", v)
		// 	fmt.Println()
		// } else {
		if !result {
			// overall = false
			// fmt.Println("Not valid JSON: ", v)
			// fmt.Println()
			return JSONWrapper{}, errors.New("invalid JSON")
		}
	}
	// fmt.Println("Overall, the JSON was: ", overall)
	return resources, nil
}

// Recursive function to unfold the nested-blocks
func printNested(t *tfjson.SchemaBlockType, depth int, attributeStr *string) {
	attLength := 0
	nestedLength := 0

	// Pre-check for ID
	if _, ok := t.Block.Attributes["id"]; ok {
		attLength++
	}

	depth++
	for k1, v1 := range t.Block.Attributes {

		// for i := 0; i < depth; i++ {
		// 	fmt.Print("\t")
		// }

		if k1 == "id" {
			ctr++
			goto RELOOP
		}

		// Attributes for top layer block
		switch v1.AttributeType.FriendlyName() {
		case "string":
			if v1.Computed {
				*attributeStr += fmt.Sprintf(attributeStrPattern, k1, "computed")
			} else if k1 == "resource_name" {
				*attributeStr += fmt.Sprintf(attributeStrPattern, k1, cfgGroupName)
			} else {
				*attributeStr += fmt.Sprintf(attributeStrPattern, k1, fmt.Sprintf("foo%d", ctr))
			}
		case "bool":
			*attributeStr += fmt.Sprintf(attributeNonStrPattern, k1, "false")
		case "number":
			*attributeStr += fmt.Sprintf(attributeNonStrPattern, k1, ctr)
		case "list of string":
			*attributeStr += "\"" + k1 + "\"" + ": ["
			*attributeStr += fmt.Sprintf("\"bar%d\", \"barbar%d\"", ctr, ctr)
			*attributeStr += "]"
		case "map of string":
			*attributeStr += "\"" + k1 + "\"" + ": {"
			*attributeStr += fmt.Sprintf("\"foo%d\": \"bar%d\", \"foofoo%d\": \"barbar%d\"", ctr, ctr, ctr, ctr)
			*attributeStr += "}"
		case "list of number":
			*attributeStr += fmt.Sprintf(attributeNonStrPattern, k1, ctr)
		case "map of number":
			*attributeStr += fmt.Sprintf(attributeNonStrPattern, k1, ctr)
		case "list of bool":
			*attributeStr += "["
			*attributeStr += fmt.Sprintf(attributeNonStrPattern, k1, "true, false, true")
			*attributeStr += "]"
		case "map of bool":
			*attributeStr += "{"
			*attributeStr += fmt.Sprintf(attributeNonStrPattern, k1, "\"thing%d\": true, \"thingy%d\": false")
			*attributeStr += "}"
		default:
			panic(fmt.Sprintf("CAUGHT EXCEPTION: %v, %v", k1, v1))
		}

		attLength++

		if (attLength < len(t.Block.Attributes)) && attLength != 0 {
			*attributeStr += ", "
		} else {
			if len(t.Block.NestedBlocks) != 0 {
				*attributeStr += ", "
			}
		}

		ctr++

	RELOOP:
		// fmt.Printf("nested block attribute: %v, schema: %v\n", k1, v1.AttributeType.FriendlyName())
	}
	for k2, v2 := range t.Block.NestedBlocks {
		// for i := 0; i < depth; i++ {
		// 	fmt.Print("\t")
		// }
		// fmt.Printf("nested nested block: %v, type: %v\n", k2, v2.NestingMode)

		if nestedLength == 0 {
			switch v2.NestingMode {
			case "single", "map", "set":
				*attributeStr += fmt.Sprintf("%q: {", k2)
			case "list":
				*attributeStr += fmt.Sprintf("%q: [{", k2)
			}
		} else {
			switch v2.NestingMode {
			case "single", "map", "set":
				*attributeStr += fmt.Sprintf(", %q: {", k2)
			case "list":
				*attributeStr += fmt.Sprintf(", %q: [{", k2)
			}
		}

		printNested(v2, depth, attributeStr)

		switch v2.NestingMode {
		case "single", "map", "set":
			*attributeStr += "}"
		case "list":
			*attributeStr += "}]"
		}

		nestedLength++
	}
}
