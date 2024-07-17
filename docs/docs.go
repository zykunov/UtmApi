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
            "name": "Igor Zykunov"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "post": {
                "description": "Создание аккаунта.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "AddAccount",
                "parameters": [
                    {
                        "description": "account add",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helpers.AccountAdd"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/helpers.AccountSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    }
                }
            }
        },
        "/accounts/{id}/balance": {
            "get": {
                "description": "Получение баланса",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "GetBalance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.BalanceSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helpers.BalanceFail"
                        }
                    }
                }
            }
        },
        "/accounts/{id}/deposit": {
            "post": {
                "description": "Внесение денег",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Deposit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "deposit",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helpers.SummAmount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.AccountSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    }
                }
            }
        },
        "/accounts/{id}/withdraw": {
            "post": {
                "description": "Снятие денег",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Withdrow",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "withdraw",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/helpers.SummAmount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.AccountSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helpers.AccountAdd": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "helpers.AccountSuccess": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "default": 500.4
                },
                "id": {
                    "type": "integer",
                    "default": 1
                },
                "name": {
                    "type": "string",
                    "default": "John"
                },
                "surname": {
                    "type": "string",
                    "default": "Doe"
                }
            }
        },
        "helpers.BalanceFail": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "default": "0"
                },
                "meta": {
                    "type": "string",
                    "default": "null"
                },
                "status_code": {
                    "type": "integer",
                    "default": 404
                }
            }
        },
        "helpers.BalanceSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "default": "400.54"
                },
                "meta": {
                    "type": "string",
                    "default": "null"
                },
                "status_code": {
                    "type": "integer",
                    "default": 200
                }
            }
        },
        "helpers.SummAmount": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "models.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "utm api",
	Description:      "REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
