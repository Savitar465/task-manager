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
        "/issues": {
            "get": {
                "description": "get all issues",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "issues"
                ],
                "summary": "Get all issues",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.IssueResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new issue using the provided IssueRequest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "issues"
                ],
                "summary": "Create an Issue",
                "parameters": [
                    {
                        "description": "Issue Request",
                        "name": "issue",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.IssueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created issue",
                        "schema": {
                            "$ref": "#/definitions/dto.IssueResponse"
                        }
                    }
                }
            }
        },
        "/issues/{id}": {
            "delete": {
                "description": "Delete an issue by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "issues"
                ],
                "summary": "Delete an Issue",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Issue ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Issue deleted"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.IssueRequest": {
            "type": "object",
            "required": [
                "assignee",
                "boardId",
                "description",
                "dueDate",
                "stageId",
                "startDate",
                "title",
                "typeId"
            ],
            "properties": {
                "assignee": {
                    "type": "string"
                },
                "boardId": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "stageId": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "typeId": {
                    "type": "string"
                }
            }
        },
        "dto.IssueResponse": {
            "type": "object",
            "properties": {
                "assignee": {
                    "type": "string"
                },
                "boardId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "stageId": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "typeId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "updatedBy": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
