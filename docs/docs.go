// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "url": "https://t.me/iivkis"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/t/answers": {
            "get": {
                "description": "Returns the created answers of the current user to the questions",
                "tags": [
                    "answers"
                ],
                "summary": "Get answers",
                "parameters": [
                    {
                        "minimum": 0,
                        "type": "integer",
                        "name": "question_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repository.AnswerModel"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "answers"
                ],
                "summary": "Create new answer for question",
                "parameters": [
                    {
                        "description": "answer body",
                        "name": "struct",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.AnswerPostBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.DefaultOut"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    }
                }
            }
        },
        "/t/answers/{id}": {
            "get": {
                "tags": [
                    "answers"
                ],
                "summary": "Get one answers by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "answer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repository.AnswerModel"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "answers"
                ],
                "summary": "Update answer fields",
                "parameters": [
                    {
                        "description": "answer body",
                        "name": "struct",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.AnswersPutBody"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "answer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.DefaultOut"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllerV1.ControllerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllerV1.AnswerPostBody": {
            "type": "object",
            "required": [
                "question_id",
                "value"
            ],
            "properties": {
                "question_id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 77
                },
                "value": {
                    "type": "string",
                    "maxLength": 1000,
                    "example": "zero"
                }
            }
        },
        "controllerV1.AnswersPutBody": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "string"
                }
            }
        },
        "controllerV1.ControllerError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "controllerV1.DefaultOut": {
            "description": "Record ID",
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "repository.AnswerModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "question_id": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0-alpha",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "ASU-Olymp API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}