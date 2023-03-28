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

func newFlowInstTask(db *gorm.DB, opts ...gen.DOOption) flowInstTask {
	_flowInstTask := flowInstTask{}

	_flowInstTask.flowInstTaskDo.UseDB(db, opts...)
	_flowInstTask.flowInstTaskDo.UseModel(&model.FlowInstTask{})

	tableName := _flowInstTask.flowInstTaskDo.TableName()
	_flowInstTask.ALL = field.NewAsterisk(tableName)
	_flowInstTask.ID = field.NewInt32(tableName, "id")
	_flowInstTask.Name = field.NewString(tableName, "name")
	_flowInstTask.FlowInstID = field.NewString(tableName, "flow_inst_id")
	_flowInstTask.FlowInstTaskID = field.NewString(tableName, "flow_inst_task_id")
	_flowInstTask.State = field.NewInt32(tableName, "state")
	_flowInstTask.CreateTime = field.NewTime(tableName, "create_time")
	_flowInstTask.CreateUser = field.NewString(tableName, "create_user")
	_flowInstTask.UpdateTime = field.NewTime(tableName, "update_time")
	_flowInstTask.UpdateUser = field.NewString(tableName, "update_user")

	_flowInstTask.fillFieldMap()

	return _flowInstTask
}

type flowInstTask struct {
	flowInstTaskDo flowInstTaskDo

	ALL            field.Asterisk
	ID             field.Int32  // 主键
	Name           field.String // 流程实例任务名称
	FlowInstID     field.String // 流程实例id
	FlowInstTaskID field.String // 流程实例任务id
	State          field.Int32  // 状态
	CreateTime     field.Time   // 创建时间
	CreateUser     field.String // 创建人
	UpdateTime     field.Time   // 更新时间
	UpdateUser     field.String // 更新人

	fieldMap map[string]field.Expr
}

func (f flowInstTask) Table(newTableName string) *flowInstTask {
	f.flowInstTaskDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f flowInstTask) As(alias string) *flowInstTask {
	f.flowInstTaskDo.DO = *(f.flowInstTaskDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *flowInstTask) updateTableName(table string) *flowInstTask {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.Name = field.NewString(table, "name")
	f.FlowInstID = field.NewString(table, "flow_inst_id")
	f.FlowInstTaskID = field.NewString(table, "flow_inst_task_id")
	f.State = field.NewInt32(table, "state")
	f.CreateTime = field.NewTime(table, "create_time")
	f.CreateUser = field.NewString(table, "create_user")
	f.UpdateTime = field.NewTime(table, "update_time")
	f.UpdateUser = field.NewString(table, "update_user")

	f.fillFieldMap()

	return f
}

func (f *flowInstTask) WithContext(ctx context.Context) *flowInstTaskDo {
	return f.flowInstTaskDo.WithContext(ctx)
}

func (f flowInstTask) TableName() string { return f.flowInstTaskDo.TableName() }

func (f flowInstTask) Alias() string { return f.flowInstTaskDo.Alias() }

func (f *flowInstTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *flowInstTask) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 9)
	f.fieldMap["id"] = f.ID
	f.fieldMap["name"] = f.Name
	f.fieldMap["flow_inst_id"] = f.FlowInstID
	f.fieldMap["flow_inst_task_id"] = f.FlowInstTaskID
	f.fieldMap["state"] = f.State
	f.fieldMap["create_time"] = f.CreateTime
	f.fieldMap["create_user"] = f.CreateUser
	f.fieldMap["update_time"] = f.UpdateTime
	f.fieldMap["update_user"] = f.UpdateUser
}

func (f flowInstTask) clone(db *gorm.DB) flowInstTask {
	f.flowInstTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f flowInstTask) replaceDB(db *gorm.DB) flowInstTask {
	f.flowInstTaskDo.ReplaceDB(db)
	return f
}

type flowInstTaskDo struct{ gen.DO }

func (f flowInstTaskDo) Debug() *flowInstTaskDo {
	return f.withDO(f.DO.Debug())
}

func (f flowInstTaskDo) WithContext(ctx context.Context) *flowInstTaskDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f flowInstTaskDo) ReadDB() *flowInstTaskDo {
	return f.Clauses(dbresolver.Read)
}

func (f flowInstTaskDo) WriteDB() *flowInstTaskDo {
	return f.Clauses(dbresolver.Write)
}

func (f flowInstTaskDo) Session(config *gorm.Session) *flowInstTaskDo {
	return f.withDO(f.DO.Session(config))
}

func (f flowInstTaskDo) Clauses(conds ...clause.Expression) *flowInstTaskDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f flowInstTaskDo) Returning(value interface{}, columns ...string) *flowInstTaskDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f flowInstTaskDo) Not(conds ...gen.Condition) *flowInstTaskDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f flowInstTaskDo) Or(conds ...gen.Condition) *flowInstTaskDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f flowInstTaskDo) Select(conds ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f flowInstTaskDo) Where(conds ...gen.Condition) *flowInstTaskDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f flowInstTaskDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *flowInstTaskDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f flowInstTaskDo) Order(conds ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f flowInstTaskDo) Distinct(cols ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f flowInstTaskDo) Omit(cols ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f flowInstTaskDo) Join(table schema.Tabler, on ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f flowInstTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f flowInstTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f flowInstTaskDo) Group(cols ...field.Expr) *flowInstTaskDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f flowInstTaskDo) Having(conds ...gen.Condition) *flowInstTaskDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f flowInstTaskDo) Limit(limit int) *flowInstTaskDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f flowInstTaskDo) Offset(offset int) *flowInstTaskDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f flowInstTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *flowInstTaskDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f flowInstTaskDo) Unscoped() *flowInstTaskDo {
	return f.withDO(f.DO.Unscoped())
}

func (f flowInstTaskDo) Create(values ...*model.FlowInstTask) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f flowInstTaskDo) CreateInBatches(values []*model.FlowInstTask, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f flowInstTaskDo) Save(values ...*model.FlowInstTask) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f flowInstTaskDo) First() (*model.FlowInstTask, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FlowInstTask), nil
	}
}

func (f flowInstTaskDo) Take() (*model.FlowInstTask, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FlowInstTask), nil
	}
}

func (f flowInstTaskDo) Last() (*model.FlowInstTask, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FlowInstTask), nil
	}
}

func (f flowInstTaskDo) Find() ([]*model.FlowInstTask, error) {
	result, err := f.DO.Find()
	return result.([]*model.FlowInstTask), err
}

func (f flowInstTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FlowInstTask, err error) {
	buf := make([]*model.FlowInstTask, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f flowInstTaskDo) FindInBatches(result *[]*model.FlowInstTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f flowInstTaskDo) Attrs(attrs ...field.AssignExpr) *flowInstTaskDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f flowInstTaskDo) Assign(attrs ...field.AssignExpr) *flowInstTaskDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f flowInstTaskDo) Joins(fields ...field.RelationField) *flowInstTaskDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f flowInstTaskDo) Preload(fields ...field.RelationField) *flowInstTaskDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f flowInstTaskDo) FirstOrInit() (*model.FlowInstTask, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FlowInstTask), nil
	}
}

func (f flowInstTaskDo) FirstOrCreate() (*model.FlowInstTask, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FlowInstTask), nil
	}
}

func (f flowInstTaskDo) FindByPage(offset int, limit int) (result []*model.FlowInstTask, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f flowInstTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f flowInstTaskDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f flowInstTaskDo) Delete(models ...*model.FlowInstTask) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *flowInstTaskDo) withDO(do gen.Dao) *flowInstTaskDo {
	f.DO = *do.(*gen.DO)
	return f
}