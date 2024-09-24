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

	"zero-fox-admin/rpc/sys/gen/model"
)

func newSysRoleMenu(db *gorm.DB, opts ...gen.DOOption) sysRoleMenu {
	_sysRoleMenu := sysRoleMenu{}

	_sysRoleMenu.sysRoleMenuDo.UseDB(db, opts...)
	_sysRoleMenu.sysRoleMenuDo.UseModel(&model.SysRoleMenu{})

	tableName := _sysRoleMenu.sysRoleMenuDo.TableName()
	_sysRoleMenu.ALL = field.NewAsterisk(tableName)
	_sysRoleMenu.ID = field.NewInt64(tableName, "id")
	_sysRoleMenu.RoleID = field.NewInt64(tableName, "role_id")
	_sysRoleMenu.MenuID = field.NewInt64(tableName, "menu_id")

	_sysRoleMenu.fillFieldMap()

	return _sysRoleMenu
}

// sysRoleMenu 角色菜单关联表
type sysRoleMenu struct {
	sysRoleMenuDo sysRoleMenuDo

	ALL    field.Asterisk
	ID     field.Int64 // 编号
	RoleID field.Int64 // 角色Id
	MenuID field.Int64 // 菜单Id

	fieldMap map[string]field.Expr
}

func (s sysRoleMenu) Table(newTableName string) *sysRoleMenu {
	s.sysRoleMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysRoleMenu) As(alias string) *sysRoleMenu {
	s.sysRoleMenuDo.DO = *(s.sysRoleMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysRoleMenu) updateTableName(table string) *sysRoleMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.RoleID = field.NewInt64(table, "role_id")
	s.MenuID = field.NewInt64(table, "menu_id")

	s.fillFieldMap()

	return s
}

func (s *sysRoleMenu) WithContext(ctx context.Context) ISysRoleMenuDo {
	return s.sysRoleMenuDo.WithContext(ctx)
}

func (s sysRoleMenu) TableName() string { return s.sysRoleMenuDo.TableName() }

func (s sysRoleMenu) Alias() string { return s.sysRoleMenuDo.Alias() }

func (s sysRoleMenu) Columns(cols ...field.Expr) gen.Columns { return s.sysRoleMenuDo.Columns(cols...) }

func (s *sysRoleMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysRoleMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["id"] = s.ID
	s.fieldMap["role_id"] = s.RoleID
	s.fieldMap["menu_id"] = s.MenuID
}

func (s sysRoleMenu) clone(db *gorm.DB) sysRoleMenu {
	s.sysRoleMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysRoleMenu) replaceDB(db *gorm.DB) sysRoleMenu {
	s.sysRoleMenuDo.ReplaceDB(db)
	return s
}

type sysRoleMenuDo struct{ gen.DO }

type ISysRoleMenuDo interface {
	gen.SubQuery
	Debug() ISysRoleMenuDo
	WithContext(ctx context.Context) ISysRoleMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysRoleMenuDo
	WriteDB() ISysRoleMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysRoleMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysRoleMenuDo
	Not(conds ...gen.Condition) ISysRoleMenuDo
	Or(conds ...gen.Condition) ISysRoleMenuDo
	Select(conds ...field.Expr) ISysRoleMenuDo
	Where(conds ...gen.Condition) ISysRoleMenuDo
	Order(conds ...field.Expr) ISysRoleMenuDo
	Distinct(cols ...field.Expr) ISysRoleMenuDo
	Omit(cols ...field.Expr) ISysRoleMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysRoleMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysRoleMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysRoleMenuDo
	Group(cols ...field.Expr) ISysRoleMenuDo
	Having(conds ...gen.Condition) ISysRoleMenuDo
	Limit(limit int) ISysRoleMenuDo
	Offset(offset int) ISysRoleMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysRoleMenuDo
	Unscoped() ISysRoleMenuDo
	Create(values ...*model.SysRoleMenu) error
	CreateInBatches(values []*model.SysRoleMenu, batchSize int) error
	Save(values ...*model.SysRoleMenu) error
	First() (*model.SysRoleMenu, error)
	Take() (*model.SysRoleMenu, error)
	Last() (*model.SysRoleMenu, error)
	Find() ([]*model.SysRoleMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRoleMenu, err error)
	FindInBatches(result *[]*model.SysRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysRoleMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysRoleMenuDo
	Assign(attrs ...field.AssignExpr) ISysRoleMenuDo
	Joins(fields ...field.RelationField) ISysRoleMenuDo
	Preload(fields ...field.RelationField) ISysRoleMenuDo
	FirstOrInit() (*model.SysRoleMenu, error)
	FirstOrCreate() (*model.SysRoleMenu, error)
	FindByPage(offset int, limit int) (result []*model.SysRoleMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysRoleMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysRoleMenuDo) Debug() ISysRoleMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysRoleMenuDo) WithContext(ctx context.Context) ISysRoleMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysRoleMenuDo) ReadDB() ISysRoleMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysRoleMenuDo) WriteDB() ISysRoleMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysRoleMenuDo) Session(config *gorm.Session) ISysRoleMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysRoleMenuDo) Clauses(conds ...clause.Expression) ISysRoleMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysRoleMenuDo) Returning(value interface{}, columns ...string) ISysRoleMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysRoleMenuDo) Not(conds ...gen.Condition) ISysRoleMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysRoleMenuDo) Or(conds ...gen.Condition) ISysRoleMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysRoleMenuDo) Select(conds ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysRoleMenuDo) Where(conds ...gen.Condition) ISysRoleMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysRoleMenuDo) Order(conds ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysRoleMenuDo) Distinct(cols ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysRoleMenuDo) Omit(cols ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysRoleMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysRoleMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysRoleMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysRoleMenuDo) Group(cols ...field.Expr) ISysRoleMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysRoleMenuDo) Having(conds ...gen.Condition) ISysRoleMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysRoleMenuDo) Limit(limit int) ISysRoleMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysRoleMenuDo) Offset(offset int) ISysRoleMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysRoleMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysRoleMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysRoleMenuDo) Unscoped() ISysRoleMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysRoleMenuDo) Create(values ...*model.SysRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysRoleMenuDo) CreateInBatches(values []*model.SysRoleMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysRoleMenuDo) Save(values ...*model.SysRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysRoleMenuDo) First() (*model.SysRoleMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) Take() (*model.SysRoleMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) Last() (*model.SysRoleMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) Find() ([]*model.SysRoleMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysRoleMenu), err
}

func (s sysRoleMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRoleMenu, err error) {
	buf := make([]*model.SysRoleMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysRoleMenuDo) FindInBatches(result *[]*model.SysRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysRoleMenuDo) Attrs(attrs ...field.AssignExpr) ISysRoleMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysRoleMenuDo) Assign(attrs ...field.AssignExpr) ISysRoleMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysRoleMenuDo) Joins(fields ...field.RelationField) ISysRoleMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysRoleMenuDo) Preload(fields ...field.RelationField) ISysRoleMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysRoleMenuDo) FirstOrInit() (*model.SysRoleMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) FirstOrCreate() (*model.SysRoleMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRoleMenu), nil
	}
}

func (s sysRoleMenuDo) FindByPage(offset int, limit int) (result []*model.SysRoleMenu, count int64, err error) {
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

func (s sysRoleMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysRoleMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysRoleMenuDo) Delete(models ...*model.SysRoleMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysRoleMenuDo) withDO(do gen.Dao) *sysRoleMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}