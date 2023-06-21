package entity

import "time"

// ModelVersionResult
// @Description: 模型版本
type ModelVersionResult struct {
	ID           int64     `json:"id"`           // 唯一id
	ModelID      string    `json:"modelID"`      // 模板id
	ModelTitle   string    `json:"modelTitle"`   // 模板版本标题
	VersionID    string    `json:"versionID"`    // 版本id
	ProcessDefID string    `json:"processDefID"` // 流程定义id
	FormDefID    string    `json:"formDefID"`    // 表单定义id
	UseStatus    int32     `json:"useStatus"`    // 使用状态【1：使用；2：未使用】
	Remark       string    `json:"remark"`       // 描述
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	CreateUser   string    `json:"createUser"`   // 创建人
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
	UpdateUser   string    `json:"updateUser"`   // 更新人
	NoticeURL    string    `json:"noticeURL"`    // 回调通知推送url
	TitleProps   string    `json:"titleProps"`   // 标题配置
}

// ModelDetailResult
// @Description: 模型详情
type ModelDetailResult struct {
	ID           int64     `json:"id"`           // 唯一id
	ModelID      string    `json:"modelID"`      // 模板id
	ModelTitle   string    `json:"modelTitle"`   // 模板标题
	ProcessDefID string    `json:"processDefID"` // 流程定义id
	FormDefID    string    `json:"formDefID"`    // 表单定义id
	ModelGroupID string    `json:"modelGroupID"` // 模版组id
	IconURL      string    `json:"iconURL"`      // icon图标地址
	Status       int32     `json:"status"`       // 模板状态【1：草稿；2：发布；3：停用】默认草稿
	Remark       string    `json:"remark"`       // 描述
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	CreateUser   string    `json:"createUser"`   // 创建人
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
	UpdateUser   string    `json:"updateUser"`   // 更新人
}