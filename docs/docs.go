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
            "email": "fiber@swagger.io"
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
        "/users/chatrooms/{guid}": {
            "get": {
                "description": "Возвращает список чатов пользователя",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Получить список чатов пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "GUID пользователя",
                        "name": "guid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ с массивом комнат",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Chatroom"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/create": {
            "post": {
                "description": "GUID нужно сохранить, нужен будет для всего",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Создание пользователя",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid input or user creation failed",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/delete/{GUID}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Удаление юзера",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/enterChatroom/{cid}/{guid}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Вход в чат",
                "parameters": [
                    {
                        "type": "string",
                        "description": "GUID пользователя",
                        "name": "guid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID чата",
                        "name": "cid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "entered",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/exitChatroom/{cid}/{guid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Выход из чата",
                "parameters": [
                    {
                        "type": "string",
                        "description": "GUID пользователя",
                        "name": "guid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID чата",
                        "name": "cid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Exited successfully",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Возвращает токен, который нужно сохранить",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Авторизация пользователя",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid input or user creation failed",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/updateEmail": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Обновление почты пользователя",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Email updated successfully",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/updatePassword": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Обновление пароля пользователя",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UpdatePasswordDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Password updated successfully",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/users/updateUsername": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Обновление юзернейма пользователя",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Username updated successfully",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        },
        "/{GUID}/{cid}": {
            "get": {
                "description": "Обновляет соединение до WebSocket'а для передачи сообщений.",
                "tags": [
                    "websocket"
                ],
                "summary": "WebSocket соединение для общения",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор пользователя",
                        "name": "GUID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор чата",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "101": {
                        "description": "Соединение обновлено до WebSocket",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверные параметры запроса",
                        "schema": {
                            "$ref": "#/definitions/entities.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Response": {
            "type": "object",
            "properties": {
                "content": {},
                "error": {
                    "type": "string"
                }
            }
        },
        "entities.UpdatePasswordDTO": {
            "type": "object",
            "properties": {
                "guid": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entities.UserDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Chatroom": {
            "type": "object",
            "properties": {
                "chatroomId": {
                    "type": "string"
                },
                "isPrivate": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "ownerGUID": {
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
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
