# Alpha Warning

This is demo code. Whilst it works for certain use cases, I'm not sure how happy it would be beyond the MVP it's designed for.

It's only been tested with OSX and the issue being, the sample provider included here is compiled for OSX. It should work with any provider in theory, but you've been warned. 

## Purpose

This project instantiates a Terraform provider (must be above version 0.13 supported) and outputs a valid JSON schema and canonical examples that have been passed through a JSON validator.

In the `/upload` directory, there is a provider ready to go.

## Setup

Install the `tfgenson` requirement on Docker.

`docker run -d --rm --name tfgenson -p 5001:5000 davedotdev/flasktfgenson`

Ensure that the entries in `config.toml` are correct for your environment.

Ready to run? Hold on to your pants Dorothy.

```bash
go mod download
go build
./terrascout -config config.toml
```

The output should be something similar to this.


```json
{
  "Schema": {
    "format_version": "0.2",
    "provider_schemas": {
      "terrascout/providers/junos-qfx": {
        "provider": {
          "version": 0,
          "block": {
            "attributes": {
              "host": {
                "type": "string",
                "description_kind": "plain",
                "required": true
              },
              "password": {
                "type": "string",
                "description_kind": "plain",
                "required": true
              },
              "port": {
                "type": "number",
                "description_kind": "plain",
                "required": true
              },
              "sshkey": {
                "type": "string",
                "description_kind": "plain",
                "required": true
              },
              "username": {
                "type": "string",
                "description_kind": "plain",
                "required": true
              }
            },
            "description_kind": "plain"
          }
        },
        "resource_schemas": {
          "junos-qfx_commit": {
            "version": 0,
            "block": {
              "attributes": {
                "id": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true,
                  "computed": true
                },
                "resource_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                }
              },
              "description_kind": "plain"
            }
          },
          "junos-qfx_inet-iface": {
            "version": 0,
            "block": {
              "attributes": {
                "commit": {
                  "type": "bool",
                  "description_kind": "plain",
                  "optional": true
                },
                "id": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true,
                  "computed": true
                },
                "iface_desc": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "iface_inet_address": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "iface_mtu": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "iface_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "iface_speed": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "iface_unit": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "resource_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                }
              },
              "description_kind": "plain"
            }
          },
          "junos-qfx_native-bgp-peer": {
            "version": 0,
            "block": {
              "attributes": {
                "bgp_group": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "bgp_local_as": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "bgp_neighbor": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "bgp_peer_as": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "bgp_peer_type": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "commit": {
                  "type": "bool",
                  "description_kind": "plain",
                  "optional": true
                },
                "id": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true,
                  "computed": true
                },
                "resource_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                }
              },
              "description_kind": "plain"
            }
          },
          "junos-qfx_vlan": {
            "version": 0,
            "block": {
              "attributes": {
                "commit": {
                  "type": "bool",
                  "description_kind": "plain",
                  "optional": true
                },
                "id": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true,
                  "computed": true
                },
                "resource_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "vlan_desc": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "vlan_l3iface": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "vlan_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "vlan_num": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                }
              },
              "description_kind": "plain"
            }
          },
          "junos-qfx_vlan-access-port": {
            "version": 0,
            "block": {
              "attributes": {
                "commit": {
                  "type": "bool",
                  "description_kind": "plain",
                  "optional": true
                },
                "id": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true,
                  "computed": true
                },
                "port_desc": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "port_duplex": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_mtu": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "port_speed": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_vlan": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "resource_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                }
              },
              "description_kind": "plain"
            }
          },
          "junos-qfx_vlan-trunk-port": {
            "version": 0,
            "block": {
              "attributes": {
                "id": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true,
                  "computed": true
                },
                "port_desc": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_duplex": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_mtu": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "port_native_vlan": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_speed": {
                  "type": "string",
                  "description_kind": "plain",
                  "optional": true
                },
                "port_vlan": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                },
                "resource_name": {
                  "type": "string",
                  "description_kind": "plain",
                  "required": true
                }
              },
              "description_kind": "plain"
            }
          }
        }
      }
    }
  },
  "Resources": [
    {
      "resource": {
        "junos-qfx_vlan-trunk-port": {
          "config-group-name": {
            "port_speed": "foo0",
            "resource_name": "config-group-name",
            "port_duplex": "foo3",
            "port_mtu": "foo4",
            "port_name": "foo5",
            "port_native_vlan": "foo6",
            "port_desc": "foo7",
            "port_vlan": "foo8"
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_commit": {
          "config-group-name": {
            "resource_name": "config-group-name"
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_inet-iface": {
          "config-group-name": {
            "commit": false,
            "iface_desc": "foo13",
            "iface_inet_address": "foo14",
            "iface_unit": "foo15",
            "resource_name": "config-group-name",
            "iface_mtu": "foo17",
            "iface_name": "foo18",
            "iface_speed": "foo19"
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_native-bgp-peer": {
          "config-group-name": {
            "resource_name": "config-group-name",
            "bgp_group": "foo21",
            "bgp_local_as": "foo22",
            "bgp_neighbor": "foo23",
            "bgp_peer_as": "foo24",
            "bgp_peer_type": "foo25",
            "commit": false
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_vlan": {
          "config-group-name": {
            "commit": false,
            "resource_name": "config-group-name",
            "vlan_desc": "foo31",
            "vlan_l3iface": "foo32",
            "vlan_name": "foo33",
            "vlan_num": "foo34"
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_vlan-access-port": {
          "config-group-name": {
            "port_duplex": "foo35",
            "port_mtu": "foo36",
            "port_vlan": "foo37",
            "resource_name": "config-group-name",
            "port_speed": "foo39",
            "commit": false,
            "port_desc": "foo42",
            "port_name": "foo43"
          }
        }
      }
    }
  ]
}
```