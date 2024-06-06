// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDictType = "sys_dict_type"

// SysDictType 字典类型表
type SysDictType struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                         // 编号
	DictName   string     `gorm:"column:dict_name;not null;comment:字典名称" json:"dictName"`                               // 字典名称
	DictType   string     `gorm:"column:dict_type;not null;comment:字典类型" json:"dictType"`                               // 字典类型
	DictStatus int32      `gorm:"column:dict_status;not null;default:1;comment:字典状态" json:"dictStatus"`                 // 字典状态
	Remark     string     `gorm:"column:remark;not null;comment:备注信息" json:"remark"`                                    // 备注信息
	IsSystem   int32      `gorm:"column:is_system;not null;comment:是否系统预留  0：否  1：是" json:"isSystem"`                   // 是否系统预留  0：否  1：是
	IsDeleted  int32      `gorm:"column:is_deleted;not null;comment:是否删除  0：否  1：是" json:"isDeleted"`                   // 是否删除  0：否  1：是
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建者" json:"createBy"`                                // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by;not null;comment:更新者" json:"updateBy"`                                // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`                                    // 更新时间
}

// TableName SysDictType's table name
func (*SysDictType) TableName() string {
	return TableNameSysDictType
}