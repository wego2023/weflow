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
                        "description": "提交审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "createTimeStart",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
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
                        "description": "返回结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bo.InstTaskResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/model/group/add": {
            "post": {
                "description": "添加模板组",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "模板"
                ],
                "summary": "添加模板组",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "ModelGroupAddVO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.ModelGroupAddVO"
                        }
                    }
                ],
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
        "/model/group/del": {
            "post": {
                "description": "删除模板组",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "模板"
                ],
                "summary": "删除模板组",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "ModelGroupDelVO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.ModelGroupDelVO"
                        }
                    }
                ],
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
        "/model/group/edit": {
            "post": {
                "description": "编辑模板组",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "模板"
                ],
                "summary": "编辑模板组",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "ModelGroupEditVO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.ModelGroupEditVO"
                        }
                    }
                ],
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
                        "description": "返回结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bo.ModelGroupResult"
                                        }
                                    }
                                }
                            ]
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
                        "description": "返回结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bo.ModelDetailResult"
                                        }
                                    }
                                }
                            ]
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
                        "description": "返回结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bo.ModelDetailResult"
                                        }
                                    }
                                }
                            ]
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
                        "description": "提交审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
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
                        "description": "完成审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
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
                        "description": "返回结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bo.UserTaskResult"
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
                        "description": "提交审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
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
                        "description": "完成审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
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
                        "description": "返回结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bo.UserTaskResult"
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
                        "description": "提交审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "createTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "提交审批时间",
                        "description": "提交审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
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
                        "description": "完成审批时间-结束 yyyy-MM-dd HH:mm:ss:SSS",
                        "name": "finishTimeEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "完成审批时间",
                        "description": "完成审批时间-开始 yyyy-MM-dd HH:mm:ss:SSS",
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
                        "description": "返回结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bo.UserTaskTodoResult"
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
                    "description": "创建时间 yyyy-MM-dd HH:mm:ss:SSS",
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
                    "description": "结束时间 yyyy-MM-dd HH:mm:ss:SSS",
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
                    "description": "发起时间 yyyy-MM-dd HH:mm:ss:SSS",
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
                    "description": "更新时间 yyyy-MM-dd HH:mm:ss:SSS",
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
        "bo.ModelDetailResult": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间 yyyy-MM-dd HH:mm:ss:SSS",
                    "type": "string"
                },
                "createUser": {
                    "description": "创建人",
                    "type": "string"
                },
                "formDefID": {
                    "description": "表单定义id",
                    "type": "string"
                },
                "iconURL": {
                    "description": "icon图标地址",
                    "type": "string"
                },
                "id": {
                    "description": "唯一id",
                    "type": "integer"
                },
                "modelGroupID": {
                    "description": "模版组id",
                    "type": "string"
                },
                "modelID": {
                    "description": "模板id",
                    "type": "string"
                },
                "modelTitle": {
                    "description": "模板标题",
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
                "status": {
                    "description": "模板状态【1：草稿；2：发布；3：停用】默认草稿",
                    "type": "integer"
                },
                "updateTime": {
                    "description": "更新时间 yyyy-MM-dd HH:mm:ss:SSS",
                    "type": "string"
                },
                "updateUser": {
                    "description": "更新人",
                    "type": "string"
                }
            }
        },
        "bo.ModelGroupResult": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间 yyyy-MM-dd HH:mm:ss:SSS",
                    "type": "string"
                },
                "createUser": {
                    "description": "创建人",
                    "type": "string"
                },
                "groupID": {
                    "description": "组id",
                    "type": "string"
                },
                "groupName": {
                    "description": "组名称",
                    "type": "string"
                },
                "id": {
                    "description": "唯一id",
                    "type": "integer"
                },
                "remark": {
                    "description": "描述",
                    "type": "string"
                },
                "updateTime": {
                    "description": "更新时间 yyyy-MM-dd HH:mm:ss:SSS",
                    "type": "string"
                },
                "updateUser": {
                    "description": "更新人",
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
                    "description": "结束时间 yyyy-MM-dd HH:mm:ss:SSS",
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
                    "description": "发起时间 yyyy-MM-dd HH:mm:ss:SSS",
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
        "vo.ModelGroupAddVO": {
            "type": "object",
            "properties": {
                "groupName": {
                    "description": "组名称",
                    "type": "string"
                },
                "remark": {
                    "description": "描述",
                    "type": "string"
                }
            }
        },
        "vo.ModelGroupDelVO": {
            "type": "object",
            "properties": {
                "groupID": {
                    "description": "组id",
                    "type": "string"
                }
            }
        },
        "vo.ModelGroupEditVO": {
            "type": "object",
            "properties": {
                "groupID": {
                    "description": "组id",
                    "type": "string"
                },
                "groupName": {
                    "description": "组名称",
                    "type": "string"
                },
                "remark": {
                    "description": "描述",
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
