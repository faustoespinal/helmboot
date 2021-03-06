package models

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qri-io/jsonschema"
)

const jsonSchema = `
{
    "$id": "helmboot-schema",
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$defs": {
        "resource_spec": {
            "type": "object",
            "properties": {
                "cpu": {
                    "type": "string"
                },
                "memory": {
                    "type": "string"
                }
            }
        },
        "env_var": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "port": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "containerPort": {
                    "type": "number"
                }
            }
        },
        "storage_mount": {
            "type": "object",
            "properties": {
                "mount": {
                    "type": "string"
                }
            }
        },
        "secret": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "pvc": {
            "type": "object",
            "properties": {
                "size": {
                    "type": "string"
                },
                "mode": {
                    "type": "string"
                },
                "storageClass": {
                    "type": "string"
                }
            },
            "required": [
                "size",
                "mode"
            ]
        },
        "app_role": {
            "type": "object",
            "properties": {
                "scopes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "app_service": {
            "type": "object",
            "properties": {
                "deployment": {
                    "type": "string"
                }
            }
        },
        "app_ingress": {
            "type": "object",
            "properties": {
                "service": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "externalService": {
                    "type": "string"
                }
            }
        },
        "app_configmap": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "container_spec": {
            "type": "object",
            "required": [
                "image",
                "tag"
            ],
            "properties": {
                "image": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "env": {
                    "type": "array",
                    "items": {
                        "$ref": "#/$defs/env_var"
                    },
                    "default": []
                },
                "configmaps": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "secrets": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ports": {
                    "type": "array",
                    "items": {
                        "$ref": "#/$defs/port"
                    },
                    "default": []
                },
                "storage": {
                    "$ref": "#/$defs/storage_mount",
                    "optional": true
                },
                "databases": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "optional": true
                },
                "messaging": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "optional": true
                },
                "resources": {
                    "type": "object",
                    "properties": {
                        "requests": {
                            "$ref": "#/$defs/resource_spec",
                            "optional": true
                        },
                        "limits": {
                            "$ref": "#/$defs/resource_spec",
                            "optional": true
                        }
                    },
                    "optional": true
                }
            }
        }
    },
    "title": "Helmboot Schema",
    "type": "object",
    "properties": {
        "apiVersion": {
            "type": "string"
        },
        "type": {
            "type": "string"
        },
        "description": {
            "type": "string"
        },
        "version": {
            "type": "string"
        },
        "appVersion": {
            "type": "string"
        },
        "spec": {
            "type": "object",
            "properties": {
                "security": {
                    "type": "object",
                    "properties": {
                        "grantTypes": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        },
                        "roles": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "patternProperties": {
                                    "^.$": {
                                        "$ref": "#/$defs/app_role"
                                    }
                                }
                            }
                        }
                    }
                },
                "deployments": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "$ref": "#/$defs/container_spec"
                            }
                        }
                    }
                },
                "jobs": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "$ref": "#/$defs/container_spec"
                            }
                        }
                    }
                },
                "services": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "$ref": "#/$defs/app_service"
                            }
                        }
                    }
                },
                "ingresses": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "$ref": "#/$defs/app_ingress"
                            }
                        }
                    }
                },
                "configmaps": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "$ref": "#/$defs/app_configmap"
                            }
                        }
                    }
                },
                "secrets": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "$ref": "#/$defs/secret"
                            }
                        }
                    }
                },
                "storage": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "patternProperties": {
                            "^.$": {
                                "$ref": "#/$defs/pvc"
                            }
                        }
                    }
                },
                "databases": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "optional": true
                },
                "messaging": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "optional": true
                },
                "testing": {
                    "type": "object",
                    "properties": {
                        "image": {
                            "type": "string"
                        },
                        "command": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        },
                        "args": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "required": [
                        "image",
                        "command",
                        "args"
                    ],
                    "optional": true
                }
            }
        }
    },
    "required": [
        "apiVersion",
        "type",
        "version",
        "spec"
    ]
}
`

// Validate JSON document per the rules defined in the json schema document.
func ValidateJson(jsonDoc []byte) []jsonschema.KeyError {
	ctx := context.Background()
	var schemaData = []byte(jsonSchema)
	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("Unmarshalling schema: " + err.Error())
	}

	errs, err := rs.ValidateBytes(ctx, jsonDoc)
	if err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
		return errs
	}
	return nil
}
