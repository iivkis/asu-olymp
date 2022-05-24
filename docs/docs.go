// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
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
        "/signIn": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Sign in user profile",
                "operationId": "SignIn",
                "parameters": [
                    {
                        "description": "sign in data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.AuthSignInBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.AuthSignInOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/signUp": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Create a new user profile",
                "operationId": "SignUp",
                "parameters": [
                    {
                        "description": "sign up data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.AuthSignUpBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.DefaultOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/t/answers": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "description": "Returns the created answers of the current user to the questions",
                "tags": [
                    "answers"
                ],
                "summary": "Get answers",
                "operationId": "GetAnswers",
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/repository.AnswerModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Create new answer for question",
                "operationId": "AddAnswer",
                "parameters": [
                    {
                        "description": "answer body",
                        "name": "body",
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.DefaultOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/t/answers/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Get one answers by ID",
                "operationId": "GetOneAnswer",
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/repository.AnswerModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Update answer fields",
                "operationId": "UpdateAnswer",
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.DefaultOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/t/questions": {
            "get": {
                "tags": [
                    "questions"
                ],
                "summary": "Get questions",
                "operationId": "GetQuestions",
                "parameters": [
                    {
                        "minimum": 0,
                        "type": "integer",
                        "name": "task_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/repository.QuestionModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Create a new question for task",
                "operationId": "AddQuestion",
                "parameters": [
                    {
                        "description": "question body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.QuestionsPostBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.DefaultOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/t/questions/{id}": {
            "get": {
                "tags": [
                    "questions"
                ],
                "summary": "Get one question by ID",
                "operationId": "GetOneQuestion",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/repository.QuestionModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Update question fields",
                "operationId": "UpdateQuestion",
                "parameters": [
                    {
                        "description": "question body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.QuestionsPutBody"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.DefaultOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/t/tasks": {
            "get": {
                "tags": [
                    "tasks"
                ],
                "summary": "Get tasks",
                "operationId": "GetTasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/repository.TasksFindResult"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create a new task",
                "operationId": "AddTask",
                "parameters": [
                    {
                        "description": "task body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.TasksPostBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.DefaultOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/t/tasks/{id}": {
            "get": {
                "tags": [
                    "tasks"
                ],
                "summary": "Get one task by ID",
                "operationId": "GetOneTask",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/repository.TasksFindResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKey": []
                    }
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update task fields",
                "operationId": "UpdateTask",
                "parameters": [
                    {
                        "description": "task body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllerV1.TasksPutBody"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controllerV1.wrap"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllerV1.DefaultOut"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
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
        "controllerV1.AuthSignInBody": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "example@mail.ru"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 4,
                    "example": "qwerty27"
                }
            }
        },
        "controllerV1.AuthSignInOut": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsImlzcyI6ImFzdS1vbHltcCJ9.NPFZIvICrpfdqUlbr5vfvRMCHgbKj28eXmLjftWPjyc"
                }
            }
        },
        "controllerV1.AuthSignUpBody": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3,
                    "example": "example@mail.ru"
                },
                "full_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1,
                    "example": "Фёдоров И.С."
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 4,
                    "example": "qwerty27"
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
        "controllerV1.QuestionsPostBody": {
            "type": "object",
            "required": [
                "text"
            ],
            "properties": {
                "task_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "text": {
                    "type": "string",
                    "maxLength": 1000
                }
            }
        },
        "controllerV1.QuestionsPutBody": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "controllerV1.TasksPostBody": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "maxLength": 2000,
                    "minLength": 10
                },
                "title": {
                    "type": "string",
                    "maxLength": 200
                }
            }
        },
        "controllerV1.TasksPutBody": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "controllerV1.wrap": {
            "description": "Standard wrapper for responses",
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "boolean"
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
        },
        "repository.QuestionModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "task_id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "repository.TasksFindResult": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer"
                },
                "author_name": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "solutions_count": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKey": {
            "description": "JWT token for authorization",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
