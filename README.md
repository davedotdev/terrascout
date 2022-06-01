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
            "iface_inet_address": "foo2",
            "iface_mtu": "foo3",
            "iface_speed": "foo4",
            "iface_unit": "foo5",
            "resource_name": "config-group-name",
            "commit": false,
            "iface_desc": "foo9",
            "iface_name": "foo10"
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_native-bgp-peer": {
          "config-group-name": {
            "resource_name": "config-group-name",
            "bgp_group": "foo12",
            "bgp_local_as": "foo13",
            "bgp_neighbor": "foo14",
            "bgp_peer_as": "foo15",
            "bgp_peer_type": "foo16",
            "commit": false
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_vlan": {
          "config-group-name": {
            "resource_name": "config-group-name",
            "vlan_desc": "foo20",
            "vlan_l3iface": "foo21",
            "vlan_name": "foo22",
            "vlan_num": "foo23",
            "commit": false
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_vlan-access-port": {
          "config-group-name": {
            "port_desc": "foo26",
            "port_mtu": "foo27",
            "port_speed": "foo28",
            "port_vlan": "foo29",
            "commit": false,
            "resource_name": "config-group-name",
            "port_duplex": "foo33",
            "port_name": "foo34"
          }
        }
      }
    },
    {
      "resource": {
        "junos-qfx_vlan-trunk-port": {
          "config-group-name": {
            "port_native_vlan": "foo35",
            "port_speed": "foo36",
            "resource_name": "config-group-name",
            "port_duplex": "foo38",
            "port_mtu": "foo39",
            "port_name": "foo40",
            "port_vlan": "foo41",
            "port_desc": "foo43"
          }
        }
      }
    }
  ],
  "JSONSchema": [
    {
      "$schema": "http://json-schema.org/schema#",
      "properties": {
        "resource": {
          "properties": {
            "junos-qfx_commit": {
              "properties": {
                "config-group-name": {
                  "properties": {
                    "resource_name": {
                      "type": "string"
                    }
                  },
                  "type": "object"
                }
              },
              "required": [
                "config-group-name"
              ],
              "type": "object"
            }
          },
          "required": [
            "junos-qfx_commit"
          ],
          "type": "object"
        }
      },
      "required": [
        "resource"
      ],
      "type": "object"
    },
    {
      "$schema": "http://json-schema.org/schema#",
      "properties": {
        "resource": {
          "properties": {
            "junos-qfx_inet-iface": {
              "properties": {
                "config-group-name": {
                  "properties": {
                    "commit": {
                      "type": "boolean"
                    },
                    "iface_desc": {
                      "type": "string"
                    },
                    "iface_inet_address": {
                      "type": "string"
                    },
                    "iface_mtu": {
                      "type": "string"
                    },
                    "iface_name": {
                      "type": "string"
                    },
                    "iface_speed": {
                      "type": "string"
                    },
                    "iface_unit": {
                      "type": "string"
                    },
                    "resource_name": {
                      "type": "string"
                    }
                  },
                  "type": "object"
                }
              },
              "required": [
                "config-group-name"
              ],
              "type": "object"
            }
          },
          "required": [
            "junos-qfx_inet-iface"
          ],
          "type": "object"
        }
      },
      "required": [
        "resource"
      ],
      "type": "object"
    },
    {
      "$schema": "http://json-schema.org/schema#",
      "properties": {
        "resource": {
          "properties": {
            "junos-qfx_native-bgp-peer": {
              "properties": {
                "config-group-name": {
                  "properties": {
                    "bgp_group": {
                      "type": "string"
                    },
                    "bgp_local_as": {
                      "type": "string"
                    },
                    "bgp_neighbor": {
                      "type": "string"
                    },
                    "bgp_peer_as": {
                      "type": "string"
                    },
                    "bgp_peer_type": {
                      "type": "string"
                    },
                    "commit": {
                      "type": "boolean"
                    },
                    "resource_name": {
                      "type": "string"
                    }
                  },
                  "type": "object"
                }
              },
              "required": [
                "config-group-name"
              ],
              "type": "object"
            }
          },
          "required": [
            "junos-qfx_native-bgp-peer"
          ],
          "type": "object"
        }
      },
      "required": [
        "resource"
      ],
      "type": "object"
    },
    {
      "$schema": "http://json-schema.org/schema#",
      "properties": {
        "resource": {
          "properties": {
            "junos-qfx_vlan": {
              "properties": {
                "config-group-name": {
                  "properties": {
                    "commit": {
                      "type": "boolean"
                    },
                    "resource_name": {
                      "type": "string"
                    },
                    "vlan_desc": {
                      "type": "string"
                    },
                    "vlan_l3iface": {
                      "type": "string"
                    },
                    "vlan_name": {
                      "type": "string"
                    },
                    "vlan_num": {
                      "type": "string"
                    }
                  },
                  "type": "object"
                }
              },
              "required": [
                "config-group-name"
              ],
              "type": "object"
            }
          },
          "required": [
            "junos-qfx_vlan"
          ],
          "type": "object"
        }
      },
      "required": [
        "resource"
      ],
      "type": "object"
    },
    {
      "$schema": "http://json-schema.org/schema#",
      "properties": {
        "resource": {
          "properties": {
            "junos-qfx_vlan-access-port": {
              "properties": {
                "config-group-name": {
                  "properties": {
                    "commit": {
                      "type": "boolean"
                    },
                    "port_desc": {
                      "type": "string"
                    },
                    "port_duplex": {
                      "type": "string"
                    },
                    "port_mtu": {
                      "type": "string"
                    },
                    "port_name": {
                      "type": "string"
                    },
                    "port_speed": {
                      "type": "string"
                    },
                    "port_vlan": {
                      "type": "string"
                    },
                    "resource_name": {
                      "type": "string"
                    }
                  },
                  "type": "object"
                }
              },
              "required": [
                "config-group-name"
              ],
              "type": "object"
            }
          },
          "required": [
            "junos-qfx_vlan-access-port"
          ],
          "type": "object"
        }
      },
      "required": [
        "resource"
      ],
      "type": "object"
    },
    {
      "$schema": "http://json-schema.org/schema#",
      "properties": {
        "resource": {
          "properties": {
            "junos-qfx_vlan-trunk-port": {
              "properties": {
                "config-group-name": {
                  "properties": {
                    "port_desc": {
                      "type": "string"
                    },
                    "port_duplex": {
                      "type": "string"
                    },
                    "port_mtu": {
                      "type": "string"
                    },
                    "port_name": {
                      "type": "string"
                    },
                    "port_native_vlan": {
                      "type": "string"
                    },
                    "port_speed": {
                      "type": "string"
                    },
                    "port_vlan": {
                      "type": "string"
                    },
                    "resource_name": {
                      "type": "string"
                    }
                  },
                  "type": "object"
                }
              },
              "required": [
                "config-group-name"
              ],
              "type": "object"
            }
          },
          "required": [
            "junos-qfx_vlan-trunk-port"
          ],
          "type": "object"
        }
      },
      "required": [
        "resource"
      ],
      "type": "object"
    }
  ]
}
```