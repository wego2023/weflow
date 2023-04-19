package entity

// StartNodeModel 开始节点
type StartNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// ApprovalNodeModel 审批节点
type ApprovalNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// NotifyNodeModel 知会节点
type NotifyNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// CustomNodeModel 自定义节点
type CustomNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// ConditionNodeModel 条件节点
type ConditionNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// BranchNodeModel 分支节点
type BranchNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// ConvergenceNodeModel 汇聚节点
type ConvergenceNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// EndNodeModel 结束节点
type EndNodeModel struct {
	NodeModel int8   `json:"nodeModel"` // 节点类型
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeId    string `json:"nodeId"`    // 节点ID
	ParentId  string `json:"parentId"`  // 父节点ID
}

// NodeModelEntity 节点类型实体
type NodeModelEntity struct {
	NodeModel      int8                `json:"nodeModel"`                // 节点类型
	NodeName       string              `json:"nodeName"`                 // 节点名称
	NodeId         string              `json:"nodeId"`                   // 节点ID
	ParentId       string              `json:"parentId"`                 // 父节点ID
	ConnData       string              `json:"connData,omitempty"`       // 连接数据
	Conditions     string              `json:"conditions,omitempty"`     // 条件
	ForwardMode    int8                `json:"forwardMode,omitempty"`    // 转发模式
	CompleteConn   int8                `json:"completeConn,omitempty"`   // 完成连接
	PermissionMode int8                `json:"permissionMode,omitempty"` // 权限模式
	AllowAdd       int8                `json:"allowAdd,omitempty"`       // 允许添加
	ProcessMode    int8                `json:"processMode,omitempty"`    // 处理模式
	TimeLimit      int64               `json:"timeLimit,omitempty"`      // 时限
	PerData        string              `json:"perData,omitempty"`        // 权限数据
	HandlerList    string              `json:"handlerList,omitempty"`    // 处理人列表
	MsgConfigList  string              `json:"msgConfigList,omitempty"`  // 消息配置列表
	Children       [][]NodeModelEntity `json:"children"`                 // 子节点
}

type NodeModelBO struct {
	NodeModel      int8       `json:"nodeModel",redis:"nodeModel"`                     // 节点类型
	NodeName       string     `json:"nodeName",redis:"nodeName"`                       // 节点名称
	NodeId         string     `json:"nodeId",redis:"nodeId"`                           // 节点ID
	ParentId       string     `json:"parentId",redis:"parentId"`                       // 父节点ID
	ConnData       string     `json:"connData,omitempty",redis:"connData"`             // 连接数据
	Conditions     string     `json:"conditions,omitempty",redis:"conditions"`         // 条件
	ForwardMode    int8       `json:"forwardMode,omitempty",redis:"forwardMode"`       // 转发模式
	CompleteConn   int8       `json:"completeConn,omitempty",redis:"completeConn"`     // 完成连接
	PermissionMode int8       `json:"permissionMode,omitempty",redis:"permissionMode"` // 权限模式
	AllowAdd       int8       `json:"allowAdd,omitempty",redis:"allowAdd"`             // 允许添加
	ProcessMode    int8       `json:"processMode,omitempty",redis:"processMode"`       // 处理模式
	TimeLimit      int64      `json:"timeLimit,omitempty",redis:"timeLimit"`           // 时限
	PerData        string     `json:"perData,omitempty",redis:"perData"`               // 权限数据
	HandlerList    string     `json:"handlerList,omitempty",redis:"handlerList"`       // 处理人列表
	MsgConfigList  string     `json:"msgConfigList,omitempty",redis:"msgConfigList"`   // 消息配置列表
	ChildrenIds    [][]string `json:"childrenIds,omitempty",redis:"childrenIds"`       // 子节点ID
	PreNodes       []string   `json:"preNodes,omitempty",redis:"preNodes"`             //上节点ID
	NextNodes      []string   `json:"nextNodes,omitempty",redis:"nextNodes"`           //下节点ID
	LastNodes      []string   `json:"lastNodes,omitempty",redis:"lastNodes"`           //分支节点尾节点ID
	Index          int        `json:"index,omitempty",redis:"index"`                   // 下标
	BranchIndex    int        `json:"branchIndex,omitempty",redis:"branchIndex"`       // 分支下标
}

//func (receiver NodeModelBO) MarshalBinary() (data []byte, err error) {
//	var buf bytes.Buffer
//	enc := gob.NewEncoder(&buf)
//	err = enc.Encode(receiver)
//	if err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//}
//
//func (receiver NodeModelBO) UnmarshalBinary(data []byte) error {
//	buf := bytes.NewBuffer(data)
//	dec := gob.NewDecoder(buf)
//	return dec.Decode(&receiver)
//}
