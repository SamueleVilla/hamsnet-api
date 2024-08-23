// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hamsters": {
            "post": {
                "description": "Responds with created hamster post",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed"
                ],
                "summary": "Responds with created hamster post",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpError"
                        }
                    }
                }
            }
        },
        "/hamsters/feed": {
            "get": {
                "description": "Responds with a list of hamster posts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed"
                ],
                "summary": "Responds with a list of hamster posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.HamsterPost"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpError"
                        }
                    }
                }
            }
        },
        "/hamsters/{id}": {
            "get": {
                "description": "Responds with the hamster post with the given id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed"
                ],
                "summary": "Responds with the hamster post with the given id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/store.HamsterPost"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpError"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Responds with a pong message",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Responds with a pong message",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HttpError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "store.HamsterPost": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "authorId": {
                    "type": "string"
                },
                "commentsCount": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageKey": {
                    "type": "string"
                },
                "likesCount": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "hamsnet.swagger.io",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Hamsnet API",
	Description:      "This is a sample server Social Network server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}