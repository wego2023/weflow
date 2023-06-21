// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "hertz-contrib",
            "url": "https://github.com/hertz-contrib"
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
        "/comm/snowflake": {
            "get": {
                "description": "获取雪花算法唯一ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共"
                ],
                "summary": "获取雪花算法唯一ID",
                "operationId": "GetSnowflake",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/insttask/initiated": {
            "get": {
                "description": "获取发起的实例任务列表（已发起）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "实例任务"
                ],
                "summary": "获取发起的实例任务列表（已发起）",
                "parameters": [
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-结束",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 格式2021-01-28 13:14:15",
                        "name": "createTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-结束",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始",
                        "name": "finishTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务状态",
                        "name": "instStatus",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "模板ID",
                        "description": "模板ID",
                        "name": "modelId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 30,
                        "description": "每页条数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "任务名称",
                        "description": "任务名称",
                        "name": "taskName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/base.Page-bo_InstTaskResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "records": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/bo.InstTaskResult"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/model/group/list": {
            "get": {
                "description": "查询获取模板组列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "模板"
                ],
                "summary": "查询获取模板组列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.Response"
                        }
                    }
                }
            }
        },
        "/model/list": {
            "get": {
                "description": "获取模板列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "模板"
                ],
                "summary": "获取模板列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.Response"
                        }
                    }
                }
            }
        },
        "/model/page": {
            "get": {
                "description": "分页获取模板列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "模板"
                ],
                "summary": "分页获取模板列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.Response"
                        }
                    }
                }
            }
        },
        "/usertask/done": {
            "get": {
                "description": "获取已办用户任务列表（已处理）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户任务"
                ],
                "summary": "获取已办用户任务列表（已处理）",
                "parameters": [
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-结束",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 格式2021-01-28 13:14:15",
                        "name": "createTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "创建人id",
                        "description": "创建人id",
                        "name": "createUserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-结束",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始",
                        "name": "finishTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务状态",
                        "name": "instStatus",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "模板ID",
                        "description": "模板ID",
                        "name": "modelId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 30,
                        "description": "每页条数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "任务名称",
                        "description": "任务名称",
                        "name": "taskName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/base.Page-bo_UserTaskResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "records": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/bo.UserTaskResult"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/usertask/received": {
            "get": {
                "description": "获取用户任务列表（我收到的）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户任务"
                ],
                "summary": "获取用户任务列表（我收到的）",
                "parameters": [
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-结束",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 格式2021-01-28 13:14:15",
                        "name": "createTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "创建人id",
                        "description": "创建人id",
                        "name": "createUserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-结束",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始",
                        "name": "finishTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务状态",
                        "name": "instStatus",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "模板ID",
                        "description": "模板ID",
                        "name": "modelId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 30,
                        "description": "每页条数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "任务名称",
                        "description": "任务名称",
                        "name": "taskName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/base.Page-bo_UserTaskResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "records": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/bo.UserTaskResult"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/usertask/todo": {
            "get": {
                "description": "获取待办用户任务列表（待处理）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户任务"
                ],
                "summary": "获取待办用户任务列表（待处理）",
                "parameters": [
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-结束",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 格式2021-01-28 13:14:15",
                        "name": "createTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "创建人id",
                        "description": "创建人id",
                        "name": "createUserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-结束",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始",
                        "name": "finishTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务状态",
                        "name": "instStatus",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "模板ID",
                        "description": "模板ID",
                        "name": "modelId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 30,
                        "description": "每页条数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "任务名称",
                        "description": "任务名称",
                        "name": "taskName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/base.Page-bo_UserTaskTodoResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "records": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/bo.UserTaskTodoResult"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.Page-bo_InstTaskResult": {
            "type": "object",
            "properties": {
                "pageNum": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页条数",
                    "type": "integer"
                },
                "records": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/bo.InstTaskResult"
                    }
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                }
            }
        },
        "base.Page-bo_UserTaskResult": {
            "type": "object",
            "properties": {
                "pageNum": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页条数",
                    "type": "integer"
                },
                "records": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/bo.UserTaskResult"
                    }
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                }
            }
        },
        "base.Page-bo_UserTaskTodoResult": {
            "type": "object",
            "properties": {
                "pageNum": {
                    "description": "页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页条数",
                    "type": "integer"
                },
                "records": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/bo.UserTaskTodoResult"
                    }
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                }
            }
        },
        "base.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "0:成功，其他：失败",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "msg": {
                    "description": "错误信息",
                    "type": "string"
                }
            }
        },
        "bo.InstTaskResult": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "createUserID": {
                    "description": "创建人id",
                    "type": "string"
                },
                "createUserName": {
                    "description": "创建人名称",
                    "type": "string"
                },
                "endTime": {
                    "description": "结束时间",
                    "type": "string"
                },
                "formDefID": {
                    "description": "表单定义id",
                    "type": "string"
                },
                "id": {
                    "description": "唯一id",
                    "type": "integer"
                },
                "instTaskID": {
                    "description": "实例任务id",
                    "type": "string"
                },
                "modelID": {
                    "description": "模板id",
                    "type": "string"
                },
                "processDefID": {
                    "description": "流程定义id",
                    "type": "string"
                },
                "remark": {
                    "description": "描述",
                    "type": "string"
                },
                "startTime": {
                    "description": "发起时间",
                    "type": "string"
                },
                "status": {
                    "description": "任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】",
                    "type": "integer"
                },
                "taskName": {
                    "description": "实例任务名称",
                    "type": "string"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "updateUserID": {
                    "description": "更新人id",
                    "type": "string"
                },
                "updateUserName": {
                    "description": "更新人名称",
                    "type": "string"
                },
                "versionID": {
                    "description": "版本id",
                    "type": "string"
                }
            }
        },
        "bo.UserTaskResult": {
            "type": "object",
            "properties": {
                "createUserID": {
                    "description": "创建人id",
                    "type": "string"
                },
                "createUserName": {
                    "description": "创建人名称",
                    "type": "string"
                },
                "endTime": {
                    "description": "结束时间",
                    "type": "string"
                },
                "instStatus": {
                    "description": "任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】",
                    "type": "integer"
                },
                "instTaskID": {
                    "description": "实例任务id",
                    "type": "string"
                },
                "nodeID": {
                    "description": "节点id",
                    "type": "string"
                },
                "nodeModel": {
                    "description": "节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】",
                    "type": "integer"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "nodeTaskID": {
                    "description": "节点任务id",
                    "type": "string"
                },
                "parentID": {
                    "description": "父节点id",
                    "type": "string"
                },
                "startTime": {
                    "description": "发起时间",
                    "type": "string"
                },
                "taskName": {
                    "description": "实例任务名称",
                    "type": "string"
                },
                "userTaskID": {
                    "description": "处理人任务id",
                    "type": "string"
                }
            }
        },
        "bo.UserTaskTodoResult": {
            "type": "object",
            "properties": {
                "createUserID": {
                    "description": "创建人id",
                    "type": "string"
                },
                "createUserName": {
                    "description": "创建人名称",
                    "type": "string"
                },
                "endTime": {
                    "description": "结束时间",
                    "type": "string"
                },
                "instStatus": {
                    "description": "任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】",
                    "type": "integer"
                },
                "instTaskID": {
                    "description": "实例任务id",
                    "type": "string"
                },
                "nodeID": {
                    "description": "节点id",
                    "type": "string"
                },
                "nodeModel": {
                    "description": "节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】",
                    "type": "integer"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "nodeTaskID": {
                    "description": "节点任务id",
                    "type": "string"
                },
                "parentID": {
                    "description": "父节点id",
                    "type": "string"
                },
                "startTime": {
                    "description": "发起时间",
                    "type": "string"
                },
                "taskName": {
                    "description": "实例任务名称",
                    "type": "string"
                },
                "userTaskID": {
                    "description": "处理人任务id",
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
	Schemes:          []string{"http"},
	Title:            "weflow",
	Description:      "weflow swagger api documention.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}