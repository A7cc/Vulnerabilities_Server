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
            "name": "Vulnerabilities_Server",
            "url": "https://github.com/A7cc/Vulnerabilities_Server"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "tags": [
                    "公共方法"
                ],
                "summary": "处理用户登录",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.LoginPassWordRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/loginout": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "处理用户注销",
                "responses": {}
            }
        },
        "/food/add": {
            "post": {
                "tags": [
                    "鉴权接口-菜品相关方法"
                ],
                "summary": "添加菜品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "接收添加菜品表单数据",
                        "name": "AddFoodRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.AddFoodRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/food/delete/{id}": {
            "delete": {
                "tags": [
                    "鉴权接口-菜品相关方法"
                ],
                "summary": "删除菜品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "菜品ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/food/detail": {
            "get": {
                "tags": [
                    "鉴权接口-菜品相关方法"
                ],
                "summary": "获取菜品详情信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "获取菜品ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/food/get": {
            "get": {
                "tags": [
                    "鉴权接口-菜品相关方法"
                ],
                "summary": "获取菜品列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/food/update": {
            "put": {
                "tags": [
                    "鉴权接口-菜品相关方法"
                ],
                "summary": "更新菜品信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "更新菜品信息",
                        "name": "UpdateFoodRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.UpdateFoodRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/food/upfoodicon": {
            "post": {
                "tags": [
                    "鉴权接口-菜品相关方法"
                ],
                "summary": "更新菜品ICON",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "更新菜品ICON",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/food/upfoodvideo": {
            "post": {
                "tags": [
                    "鉴权接口-菜品相关方法"
                ],
                "summary": "更新菜品Video",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "更新菜品视频",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/home/get": {
            "get": {
                "tags": [
                    "鉴权接口-首页设置方法"
                ],
                "summary": "获取系统数据信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/home/getsentence": {
            "get": {
                "tags": [
                    "鉴权接口-首页设置方法"
                ],
                "summary": "获取每日金句",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "获取每日金句的url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/home/updateInfo": {
            "put": {
                "tags": [
                    "鉴权接口-首页设置方法"
                ],
                "summary": "用户自己更新个人信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "用户自己更新个人信息参数",
                        "name": "UpdateInfoType",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.UpdateInfoType"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/home/updatePwd": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "鉴权接口-首页设置方法"
                ],
                "summary": "修改个人密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "用户自己更新密码和UID",
                        "name": "newpwdinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/home/upuseravatar": {
            "post": {
                "tags": [
                    "鉴权接口-首页设置方法"
                ],
                "summary": "上传头像",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "上传头像",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/order/add": {
            "post": {
                "tags": [
                    "鉴权接口-订单相关方法"
                ],
                "summary": "新增订单信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "添加订单信息",
                        "name": "AddOrderRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.AddOrderRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/order/delete/{id}": {
            "delete": {
                "tags": [
                    "鉴权接口-订单相关方法"
                ],
                "summary": "删除订单信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "删除订单ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/order/detail": {
            "get": {
                "tags": [
                    "鉴权接口-订单相关方法"
                ],
                "summary": "根据ID获取订单信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "获取订单ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/order/get": {
            "get": {
                "tags": [
                    "鉴权接口-订单相关方法"
                ],
                "summary": "获取订单列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "获取订单查询参数",
                        "name": "GetOrderListRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.GetOrderListRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/role/add": {
            "post": {
                "tags": [
                    "鉴权接口-角色相关方法"
                ],
                "summary": "新增角色信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "添加角色信息",
                        "name": "AddRoleRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.AddRoleRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/role/delete/{id}": {
            "delete": {
                "tags": [
                    "鉴权接口-角色相关方法"
                ],
                "summary": "删除角色信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "删除角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/role/detail": {
            "get": {
                "tags": [
                    "鉴权接口-角色相关方法"
                ],
                "summary": "获取角色详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "获取角色ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/role/get": {
            "get": {
                "tags": [
                    "鉴权接口-角色相关方法"
                ],
                "summary": "获取角色列表信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "获取角色查询参数",
                        "name": "GetRoleListRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.GetRoleListRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/role/update": {
            "put": {
                "tags": [
                    "鉴权接口-角色相关方法"
                ],
                "summary": "修改角色信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "修改角色信息参数",
                        "name": "UpdateRoleRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.UpdateRoleRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/settings/backupsdb": {
            "get": {
                "tags": [
                    "鉴权接口-系统设置相关方法"
                ],
                "summary": "备份数据库",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/settings/deletedb": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "鉴权接口-系统设置相关方法"
                ],
                "summary": "删除备份数据库",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "删除数据库备份文件",
                        "name": "dbfile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/settings/downdb": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "鉴权接口-系统设置相关方法"
                ],
                "summary": "数据库下载",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "下载数据库备份文件",
                        "name": "dbfile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/settings/getdb": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "鉴权接口-系统设置相关方法"
                ],
                "summary": "获取备份数据库列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "数据库备份目录",
                        "name": "dir",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/settings/ping": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "鉴权接口-系统设置相关方法"
                ],
                "summary": "测试连通性",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "IP地址",
                        "name": "addre",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/add": {
            "post": {
                "tags": [
                    "鉴权接口-用户相关方法"
                ],
                "summary": "新增用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "新增用户信息",
                        "name": "AddUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.AddUserRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/delete/{id}": {
            "delete": {
                "tags": [
                    "鉴权接口-用户相关方法"
                ],
                "summary": "删除用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "删除用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/detail": {
            "get": {
                "tags": [
                    "鉴权接口-用户相关方法"
                ],
                "summary": "获取管理员详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "获取用户ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/get": {
            "get": {
                "tags": [
                    "鉴权接口-用户相关方法"
                ],
                "summary": "获取用户列表信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "获取用户查询参数",
                        "name": "GetUserListRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.GetUserListRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/update": {
            "put": {
                "tags": [
                    "鉴权接口-用户相关方法"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "更新用户信息参数",
                        "name": "UpdateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "define.AddFoodRequest": {
            "type": "object",
            "properties": {
                "foodicon": {
                    "type": "string"
                },
                "foodname": {
                    "type": "string"
                },
                "foodprocedure": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "remarks": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "video": {
                    "type": "string"
                }
            }
        },
        "define.AddOrderRequest": {
            "type": "object",
            "properties": {
                "food": {
                    "type": "string"
                },
                "num": {
                    "type": "integer"
                },
                "remarks": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "define.AddRoleRequest": {
            "type": "object",
            "properties": {
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                }
            }
        },
        "define.AddUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                },
                "sex": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "define.GetOrderListRequest": {
            "type": "object",
            "properties": {
                "keyword": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "define.GetRoleListRequest": {
            "type": "object",
            "properties": {
                "keyword": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "define.GetUserListRequest": {
            "type": "object",
            "properties": {
                "keyword": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "define.LoginPassWordRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "登录验证码",
                    "type": "string"
                },
                "password": {
                    "description": "登录密码",
                    "type": "string"
                },
                "username": {
                    "description": "登录用户名",
                    "type": "string"
                }
            }
        },
        "define.UpdateFoodRequest": {
            "type": "object",
            "properties": {
                "foodicon": {
                    "type": "string"
                },
                "foodname": {
                    "type": "string"
                },
                "foodprocedure": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "remarks": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "video": {
                    "type": "string"
                }
            }
        },
        "define.UpdateInfoType": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "sex": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "define.UpdateRoleRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                }
            }
        },
        "define.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "remarks": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                },
                "sex": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "食谱菜单管理系统靶场",
	Description:      "这是一个用Golang写的Web靶场，该系统是以食谱菜单管理系统为场景去编写，一种实战化形式的安全漏洞靶场，其中存在多个安全漏洞，需要我们去探索和发现。该项目旨在帮助安全研究人员和爱好者了解和掌握关于Golang系统的渗透测试和代码审计知识。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}