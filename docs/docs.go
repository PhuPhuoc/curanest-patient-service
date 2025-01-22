// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/patients": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create patient profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patient"
                ],
                "summary": "create patient profile",
                "parameters": [
                    {
                        "description": "account creation data",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/patientcommands.CreatePatientProfileCmdDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/api/v1/patients/relatives": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all patients belong to relatives user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "patient"
                ],
                "summary": "get all patients belong to relatives user",
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/api/v1/relatives": {
            "post": {
                "description": "create relatives account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "relative"
                ],
                "summary": "create relatives account",
                "parameters": [
                    {
                        "description": "account creation data",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/relativescommands.CreateRelativeAccountCmdDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "ping server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping server",
                "responses": {
                    "200": {
                        "description": "message success",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "patientcommands.CreatePatientProfileCmdDTO": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "desc-pathology": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full-name": {
                    "type": "string"
                },
                "note-for-nurse": {
                    "type": "string"
                },
                "phone-number": {
                    "type": "string"
                },
                "ward": {
                    "type": "string"
                }
            }
        },
        "relativescommands.CreateRelativeAccountCmdDTO": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full-name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone-number": {
                    "type": "string"
                },
                "ward": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Patient Service",
	Description:      "Auth-service: https://api.curanest.com.vn/auth/swagger/index.html.\nPatient-service: https://api.curanest.com.vn/auth/swagger/index.html.\nNurse-service: https://api.curanest.com.vn/auth/swagger/index.html.\nAppointment-service (not ready - expected): https://api.curanest.com.vn/auth/swagger/index.html.\nNotification-service (not ready - expected): https://api.curanest.com.vn/auth/swagger/index.html.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
