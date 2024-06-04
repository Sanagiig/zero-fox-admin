// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"zero-fox-admin/rpc/gen/model"
)

func newSysMenu(db *gorm.DB, opts ...gen.DOOption) sysMenu {
	_sysMenu := sysMenu{}

	_sysMenu.sysMenuDo.UseDB(db, opts...)
	_sysMenu.sysMenuDo.UseModel(&model.SysMenu{})

	tableName := _sysMenu.sysMenuDo.TableName()
	_sysMenu.ALL = field.NewAsterisk(tableName)
	_sysMenu.ID = field.NewInt64(tableName, "id")
	_sysMenu.Name = field.NewString(tableName, "name")
	_sysMenu.ParentID = field.NewInt64(tableName, "parent_id")
	_sysMenu.URL = field.NewString(tableName, "url")
	_sysMenu.Perms = field.NewString(tableName, "perms")
	_sysMenu.Type = field.NewInt32(tableName, "type")
	_sysMenu.Icon = field.NewString(tableName, "icon")
	_sysMenu.OrderNum = field.NewInt32(tableName, "order_num")
	_sysMenu.CreateBy = field.NewString(tableName, "create_by")
	_sysMenu.CreateTime = field.NewTime(tableName, "create_time")
	_sysMenu.UpdateBy = field.NewString(tableName, "update_by")
	_sysMenu.UpdateTime = field.NewTime(tableName, "update_time")
	_sysMenu.DelFlag = field.NewInt32(tableName, "del_flag")
	_sysMenu.VuePath = field.NewString(tableName, "vue_path")
	_sysMenu.VueComponent = field.NewString(tableName, "vue_component")
	_sysMenu.VueIcon = field.NewString(tableName, "vue_icon")
	_sysMenu.VueRedirect = field.NewString(tableName, "vue_redirect")
	_sysMenu.BackgroundURL = field.NewString(tableName, "background_url")

	_sysMenu.fillFieldMap()

	return _sysMenu
}

// sysMenu 菜单管理
type sysMenu struct {
	sysMenuDo sysMenuDo

	ALL           field.Asterisk
	ID            field.Int64  // 编号
	Name          field.String // 菜单名称
	ParentID      field.Int64  // 父菜单ID，一级菜单为0
	URL           field.String
	Perms         field.String
	Type          field.Int32  // 类型 0：目录,1：菜单,2：按钮,3：外链
	Icon          field.String // 菜单图标
	OrderNum      field.Int32  // 排序
	CreateBy      field.String // 创建者
	CreateTime    field.Time   // 创建时间
	UpdateBy      field.String // 更新者
	UpdateTime    field.Time   // 更新时间
	DelFlag       field.Int32  // 是否删除  0：已删除  1：正常
	VuePath       field.String // vue系统的path
	VueComponent  field.String // vue的页面
	VueIcon       field.String // vue的图标
	VueRedirect   field.String // vue的路由重定向
	BackgroundURL field.String // 后台地址

	fieldMap map[string]field.Expr
}

func (s sysMenu) Table(newTableName string) *sysMenu {
	s.sysMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysMenu) As(alias string) *sysMenu {
	s.sysMenuDo.DO = *(s.sysMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysMenu) updateTableName(table string) *sysMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.ParentID = field.NewInt64(table, "parent_id")
	s.URL = field.NewString(table, "url")
	s.Perms = field.NewString(table, "perms")
	s.Type = field.NewInt32(table, "type")
	s.Icon = field.NewString(table, "icon")
	s.OrderNum = field.NewInt32(table, "order_num")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.DelFlag = field.NewInt32(table, "del_flag")
	s.VuePath = field.NewString(table, "vue_path")
	s.VueComponent = field.NewString(table, "vue_component")
	s.VueIcon = field.NewString(table, "vue_icon")
	s.VueRedirect = field.NewString(table, "vue_redirect")
	s.BackgroundURL = field.NewString(table, "background_url")

	s.fillFieldMap()

	return s
}

func (s *sysMenu) WithContext(ctx context.Context) ISysMenuDo { return s.sysMenuDo.WithContext(ctx) }

func (s sysMenu) TableName() string { return s.sysMenuDo.TableName() }

func (s sysMenu) Alias() string { return s.sysMenuDo.Alias() }

func (s sysMenu) Columns(cols ...field.Expr) gen.Columns { return s.sysMenuDo.Columns(cols...) }

func (s *sysMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 18)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["parent_id"] = s.ParentID
	s.fieldMap["url"] = s.URL
	s.fieldMap["perms"] = s.Perms
	s.fieldMap["type"] = s.Type
	s.fieldMap["icon"] = s.Icon
	s.fieldMap["order_num"] = s.OrderNum
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["del_flag"] = s.DelFlag
	s.fieldMap["vue_path"] = s.VuePath
	s.fieldMap["vue_component"] = s.VueComponent
	s.fieldMap["vue_icon"] = s.VueIcon
	s.fieldMap["vue_redirect"] = s.VueRedirect
	s.fieldMap["background_url"] = s.BackgroundURL
}

func (s sysMenu) clone(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysMenu) replaceDB(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceDB(db)
	return s
}

type sysMenuDo struct{ gen.DO }

type ISysMenuDo interface {
	gen.SubQuery
	Debug() ISysMenuDo
	WithContext(ctx context.Context) ISysMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysMenuDo
	WriteDB() ISysMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysMenuDo
	Not(conds ...gen.Condition) ISysMenuDo
	Or(conds ...gen.Condition) ISysMenuDo
	Select(conds ...field.Expr) ISysMenuDo
	Where(conds ...gen.Condition) ISysMenuDo
	Order(conds ...field.Expr) ISysMenuDo
	Distinct(cols ...field.Expr) ISysMenuDo
	Omit(cols ...field.Expr) ISysMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	Group(cols ...field.Expr) ISysMenuDo
	Having(conds ...gen.Condition) ISysMenuDo
	Limit(limit int) ISysMenuDo
	Offset(offset int) ISysMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo
	Unscoped() ISysMenuDo
	Create(values ...*model.SysMenu) error
	CreateInBatches(values []*model.SysMenu, batchSize int) error
	Save(values ...*model.SysMenu) error
	First() (*model.SysMenu, error)
	Take() (*model.SysMenu, error)
	Last() (*model.SysMenu, error)
	Find() ([]*model.SysMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error)
	FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysMenuDo
	Assign(attrs ...field.AssignExpr) ISysMenuDo
	Joins(fields ...field.RelationField) ISysMenuDo
	Preload(fields ...field.RelationField) ISysMenuDo
	FirstOrInit() (*model.SysMenu, error)
	FirstOrCreate() (*model.SysMenu, error)
	FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysMenuDo) Debug() ISysMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysMenuDo) WithContext(ctx context.Context) ISysMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysMenuDo) ReadDB() ISysMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysMenuDo) WriteDB() ISysMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysMenuDo) Session(config *gorm.Session) ISysMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysMenuDo) Clauses(conds ...clause.Expression) ISysMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysMenuDo) Returning(value interface{}, columns ...string) ISysMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysMenuDo) Not(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysMenuDo) Or(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysMenuDo) Select(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysMenuDo) Where(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysMenuDo) Order(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysMenuDo) Distinct(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysMenuDo) Omit(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysMenuDo) Group(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysMenuDo) Having(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysMenuDo) Limit(limit int) ISysMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysMenuDo) Offset(offset int) ISysMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysMenuDo) Unscoped() ISysMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysMenuDo) Create(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysMenuDo) CreateInBatches(values []*model.SysMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysMenuDo) Save(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysMenuDo) First() (*model.SysMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Take() (*model.SysMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Last() (*model.SysMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Find() ([]*model.SysMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysMenu), err
}

func (s sysMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error) {
	buf := make([]*model.SysMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysMenuDo) FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysMenuDo) Attrs(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysMenuDo) Assign(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysMenuDo) Joins(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysMenuDo) Preload(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysMenuDo) FirstOrInit() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FirstOrCreate() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysMenuDo) Delete(models ...*model.SysMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysMenuDo) withDO(do gen.Dao) *sysMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
