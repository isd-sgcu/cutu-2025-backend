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
        "/api/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a list of all users",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch users",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a user by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Account Info",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update user role",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/qr/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a QR code URL for a user",
                "produces": [
                    "application/json"
                ],
                "summary": "Get QR code URL",
                "parameters": [
                    {
                        "type": "string",
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
                            "$ref": "#/definitions/domain.QrResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch user",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a user by its ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Scan QR code",
                "parameters": [
                    {
                        "type": "string",
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
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "User has already entered",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/register": {
            "post": {
                "description": "Register a new user in the system",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Phone",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User University",
                        "name": "university",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Jersey Size",
                        "name": "sizeJersey",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Food Limitation",
                        "name": "foodLimitation",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Invitation Code",
                        "name": "invitationCode",
                        "in": "formData"
                    },
                    {
                        "enum": [
                            "student",
                            "alumni",
                            "general_public"
                        ],
                        "type": "string",
                        "description": "User Status",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "User Image",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Graduated Year",
                        "name": "graduatedYear",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Faculty",
                        "name": "faculty",
                        "in": "formData"
                    },
                    {
                        "enum": [
                            "studying",
                            "graduated"
                        ],
                        "type": "string",
                        "description": "Education",
                        "name": "education",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create user",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/role/{id}": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    },
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a user by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update user role by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Role",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update user role",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/signin": {
            "post": {
                "description": "SignIn",
                "produces": [
                    "application/json"
                ],
                "summary": "SignIn",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to signin",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a user by its ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
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
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch user",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a user by its ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to delete user",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a user by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update user",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Education": {
            "type": "string",
            "enum": [
                "studying",
                "graduated"
            ],
            "x-enum-varnames": [
                "EducationStudying",
                "EducationGraduated"
            ]
        },
        "domain.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "domain.QrResponse": {
            "type": "object",
            "properties": {
                "qrUrl": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "domain.Role": {
            "type": "string",
            "enum": [
                "member",
                "staff",
                "admin"
            ],
            "x-enum-varnames": [
                "Member",
                "Staff",
                "Admin"
            ]
        },
        "domain.Status": {
            "type": "string",
            "enum": [
                "student",
                "alumni",
                "general_public"
            ],
            "x-enum-varnames": [
                "StatusStudent",
                "StatusAlumni",
                "StatusGeneralPublic"
            ]
        },
        "domain.TokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "education": {
                    "$ref": "#/definitions/domain.Education"
                },
                "email": {
                    "type": "string"
                },
                "faculty": {
                    "type": "string"
                },
                "foodLimitation": {
                    "type": "string"
                },
                "graduatedYear": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageURL": {
                    "type": "string"
                },
                "invitationCode": {
                    "type": "string"
                },
                "lastEntered": {
                    "description": "Timestamp for the last QR scan",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "description": "Make phone unique",
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.Role"
                },
                "sizeJersey": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.Status"
                },
                "university": {
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
