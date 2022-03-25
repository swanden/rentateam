// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/": {
            "get": {
                "description": "Show API info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api info"
                ],
                "summary": "Show API info",
                "operationId": "index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.apiInfo"
                        }
                    }
                }
            }
        },
        "/post": {
            "get": {
                "description": "Show all posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Show all posts",
                "operationId": "all",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.allResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.responseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create post",
                "operationId": "create",
                "parameters": [
                    {
                        "description": "Post fields",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.createRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.createResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.responseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.responseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.Post": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "v1.allResponse": {
            "type": "object",
            "properties": {
                "posts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Post"
                    }
                }
            }
        },
        "v1.apiInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "v1.createRequest": {
            "type": "object",
            "required": [
                "body",
                "created_at",
                "tags",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "example": "Post body"
                },
                "created_at": {
                    "type": "string",
                    "example": "2021-12-01 15:04:05"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "FirstTag",
                        "SecondTag"
                    ]
                },
                "title": {
                    "type": "string",
                    "example": "Post title"
                }
            }
        },
        "v1.createResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.responseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Rentateam blog",
	Description:      "Rentateam blog test task",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
