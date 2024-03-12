// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "url": "https://www.Panth3r.io/"
        },
        "version": "{{.Version}}",
        "x-logo": {
            "altText": "example logo",
            "backgroundColor": "#000000",
            "href": "https://example.com/img.png",
            "url": "https://example.com/img.png"
        }
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "return application status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "status"
                ],
                "summary": "Report application status",
                "responses": {
                    "200": {
                        "description": "Server_status:available",
                        "schema": {
                            "$ref": "#/definitions/http.ServerStatus"
                        }
                    },
                    "500": {
                        "description": "INTERNAL SERVRER ERROR",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sendmail": {
            "post": {
                "description": "sends user registration email",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "application"
                ],
                "summary": "sends user registration email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Users any preferred name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "email",
                        "description": "valid email address",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "CLIENT ERROR: BAD REQUEST",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "SERVER ERROR: INTERNAL SERVRER ERROR",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.ServerStatus": {
            "type": "object",
            "properties": {
                "application_Env": {
                    "type": "string"
                },
                "application_Version": {
                    "type": "string"
                },
                "server_status": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Panth3r waitlist-backend",
	Description:      "Panth3r waitlist-backend API endpoints.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
