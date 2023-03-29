// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"wego2023/weflow/pkg/model"
)

func newInstHandlerTaskOpinion(db *gorm.DB, opts ...gen.DOOption) instHandlerTaskOpinion {
	_instHandlerTaskOpinion := instHandlerTaskOpinion{}

	_instHandlerTaskOpinion.instHandlerTaskOpinionDo.UseDB(db, opts...)
	_instHandlerTaskOpinion.instHandlerTaskOpinionDo.UseModel(&model.InstHandlerTaskOpinion{})

	tableName := _instHandlerTaskOpinion.instHandlerTaskOpinionDo.TableName()
	_instHandlerTaskOpinion.ALL = field.NewAsterisk(tableName)
	_instHandlerTaskOpinion.ID = field.NewInt64(tableName, "id")
	_instHandlerTaskOpinion.InstTaskID = field.NewString(tableName, "inst_task_id")
	_instHandlerTaskOpinion.NodeTaskID = field.NewString(tableName, "node_task_id")
	_instHandlerTaskOpinion.NodeID = field.NewString(tableName, "node_id")
	_instHandlerTaskOpinion.OpinionID = field.NewString(tableName, "opinion_id")
	_instHandlerTaskOpinion.Opinion = field.NewInt32(tableName, "opinion")
	_instHandlerTaskOpinion.OpinionDesc = field.NewString(tableName, "opinion_desc")
	_instHandlerTaskOpinion.OpUserID = field.NewString(tableName, "op_user_id")
	_instHandlerTaskOpinion.OpUserName = field.NewString(tableName, "op_user_name")
	_instHandlerTaskOpinion.CreateTime = field.NewTime(tableName, "create_time")
	_instHandlerTaskOpinion.UpdateTime = field.NewTime(tableName, "update_time")
	_instHandlerTaskOpinion.OpinionTime = field.NewTime(tableName, "opinion_time")

	_instHandlerTaskOpinion.fillFieldMap()

	return _instHandlerTaskOpinion
}

type instHandlerTaskOpinion struct {
	instHandlerTaskOpinionDo instHandlerTaskOpinionDo

	ALL         field.Asterisk
	ID          field.Int64  // 唯一id
	InstTaskID  field.String // 实例任务id
	NodeTaskID  field.String // 节点任务id
	NodeID      field.String // 节点id
	OpinionID   field.String // 意见id
	Opinion     field.Int32  // 处理意见【1：未处理；2：已阅；3：同意；4：不同意；5：回退；6：终止】
	OpinionDesc field.String // 处理意见描述
	OpUserID    field.String // 操作用户id
	OpUserName  field.String // 操作用户名称
	CreateTime  field.Time   // 创建时间
	UpdateTime  field.Time   // 更新时间
	OpinionTime field.Time   // 发表意见时间

	fieldMap map[string]field.Expr
}

func (i instHandlerTaskOpinion) Table(newTableName string) *instHandlerTaskOpinion {
	i.instHandlerTaskOpinionDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i instHandlerTaskOpinion) As(alias string) *instHandlerTaskOpinion {
	i.instHandlerTaskOpinionDo.DO = *(i.instHandlerTaskOpinionDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *instHandlerTaskOpinion) updateTableName(table string) *instHandlerTaskOpinion {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.InstTaskID = field.NewString(table, "inst_task_id")
	i.NodeTaskID = field.NewString(table, "node_task_id")
	i.NodeID = field.NewString(table, "node_id")
	i.OpinionID = field.NewString(table, "opinion_id")
	i.Opinion = field.NewInt32(table, "opinion")
	i.OpinionDesc = field.NewString(table, "opinion_desc")
	i.OpUserID = field.NewString(table, "op_user_id")
	i.OpUserName = field.NewString(table, "op_user_name")
	i.CreateTime = field.NewTime(table, "create_time")
	i.UpdateTime = field.NewTime(table, "update_time")
	i.OpinionTime = field.NewTime(table, "opinion_time")

	i.fillFieldMap()

	return i
}

func (i *instHandlerTaskOpinion) WithContext(ctx context.Context) *instHandlerTaskOpinionDo {
	return i.instHandlerTaskOpinionDo.WithContext(ctx)
}

func (i instHandlerTaskOpinion) TableName() string { return i.instHandlerTaskOpinionDo.TableName() }

func (i instHandlerTaskOpinion) Alias() string { return i.instHandlerTaskOpinionDo.Alias() }

func (i *instHandlerTaskOpinion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *instHandlerTaskOpinion) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 12)
	i.fieldMap["id"] = i.ID
	i.fieldMap["inst_task_id"] = i.InstTaskID
	i.fieldMap["node_task_id"] = i.NodeTaskID
	i.fieldMap["node_id"] = i.NodeID
	i.fieldMap["opinion_id"] = i.OpinionID
	i.fieldMap["opinion"] = i.Opinion
	i.fieldMap["opinion_desc"] = i.OpinionDesc
	i.fieldMap["op_user_id"] = i.OpUserID
	i.fieldMap["op_user_name"] = i.OpUserName
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["update_time"] = i.UpdateTime
	i.fieldMap["opinion_time"] = i.OpinionTime
}

func (i instHandlerTaskOpinion) clone(db *gorm.DB) instHandlerTaskOpinion {
	i.instHandlerTaskOpinionDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i instHandlerTaskOpinion) replaceDB(db *gorm.DB) instHandlerTaskOpinion {
	i.instHandlerTaskOpinionDo.ReplaceDB(db)
	return i
}

type instHandlerTaskOpinionDo struct{ gen.DO }

func (i instHandlerTaskOpinionDo) Debug() *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Debug())
}

func (i instHandlerTaskOpinionDo) WithContext(ctx context.Context) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i instHandlerTaskOpinionDo) ReadDB() *instHandlerTaskOpinionDo {
	return i.Clauses(dbresolver.Read)
}

func (i instHandlerTaskOpinionDo) WriteDB() *instHandlerTaskOpinionDo {
	return i.Clauses(dbresolver.Write)
}

func (i instHandlerTaskOpinionDo) Session(config *gorm.Session) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Session(config))
}

func (i instHandlerTaskOpinionDo) Clauses(conds ...clause.Expression) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i instHandlerTaskOpinionDo) Returning(value interface{}, columns ...string) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i instHandlerTaskOpinionDo) Not(conds ...gen.Condition) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i instHandlerTaskOpinionDo) Or(conds ...gen.Condition) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i instHandlerTaskOpinionDo) Select(conds ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i instHandlerTaskOpinionDo) Where(conds ...gen.Condition) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i instHandlerTaskOpinionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *instHandlerTaskOpinionDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i instHandlerTaskOpinionDo) Order(conds ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i instHandlerTaskOpinionDo) Distinct(cols ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i instHandlerTaskOpinionDo) Omit(cols ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i instHandlerTaskOpinionDo) Join(table schema.Tabler, on ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i instHandlerTaskOpinionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i instHandlerTaskOpinionDo) RightJoin(table schema.Tabler, on ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i instHandlerTaskOpinionDo) Group(cols ...field.Expr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i instHandlerTaskOpinionDo) Having(conds ...gen.Condition) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i instHandlerTaskOpinionDo) Limit(limit int) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i instHandlerTaskOpinionDo) Offset(offset int) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i instHandlerTaskOpinionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i instHandlerTaskOpinionDo) Unscoped() *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Unscoped())
}

func (i instHandlerTaskOpinionDo) Create(values ...*model.InstHandlerTaskOpinion) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i instHandlerTaskOpinionDo) CreateInBatches(values []*model.InstHandlerTaskOpinion, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i instHandlerTaskOpinionDo) Save(values ...*model.InstHandlerTaskOpinion) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i instHandlerTaskOpinionDo) First() (*model.InstHandlerTaskOpinion, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTaskOpinion), nil
	}
}

func (i instHandlerTaskOpinionDo) Take() (*model.InstHandlerTaskOpinion, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTaskOpinion), nil
	}
}

func (i instHandlerTaskOpinionDo) Last() (*model.InstHandlerTaskOpinion, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTaskOpinion), nil
	}
}

func (i instHandlerTaskOpinionDo) Find() ([]*model.InstHandlerTaskOpinion, error) {
	result, err := i.DO.Find()
	return result.([]*model.InstHandlerTaskOpinion), err
}

func (i instHandlerTaskOpinionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InstHandlerTaskOpinion, err error) {
	buf := make([]*model.InstHandlerTaskOpinion, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i instHandlerTaskOpinionDo) FindInBatches(result *[]*model.InstHandlerTaskOpinion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i instHandlerTaskOpinionDo) Attrs(attrs ...field.AssignExpr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i instHandlerTaskOpinionDo) Assign(attrs ...field.AssignExpr) *instHandlerTaskOpinionDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i instHandlerTaskOpinionDo) Joins(fields ...field.RelationField) *instHandlerTaskOpinionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i instHandlerTaskOpinionDo) Preload(fields ...field.RelationField) *instHandlerTaskOpinionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i instHandlerTaskOpinionDo) FirstOrInit() (*model.InstHandlerTaskOpinion, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTaskOpinion), nil
	}
}

func (i instHandlerTaskOpinionDo) FirstOrCreate() (*model.InstHandlerTaskOpinion, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTaskOpinion), nil
	}
}

func (i instHandlerTaskOpinionDo) FindByPage(offset int, limit int) (result []*model.InstHandlerTaskOpinion, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i instHandlerTaskOpinionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i instHandlerTaskOpinionDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i instHandlerTaskOpinionDo) Delete(models ...*model.InstHandlerTaskOpinion) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *instHandlerTaskOpinionDo) withDO(do gen.Dao) *instHandlerTaskOpinionDo {
	i.DO = *do.(*gen.DO)
	return i
}