// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/aircrafts": {
            "get": {
                "description": "Get aircraft",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aircrafts"
                ],
                "summary": "Get aircraft",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Aircraft"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.Aircraft": {
            "type": "object",
            "properties": {
                "aircraft_name": {
                    "type": "string"
                },
                "created_at": {
                    "$ref": "#/definitions/structs.CustomTime"
                },
                "iata_code": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "plane_type_id": {
                    "type": "string",
                    "example": "0"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "structs.CustomTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
