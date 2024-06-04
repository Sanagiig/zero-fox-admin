// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDictItem = "sys_dict_item"

// SysDictItem 字典项表
type SysDictItem struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                         // 编号
	DictType   string     `gorm:"column:dict_type;not null;comment:字典类型" json:"dictType"`                               // 字典类型
	DictLabel  string     `gorm:"column:dict_label;not null;comment:字典标签" json:"dictLabel"`                             // 字典标签
	DictValue  string     `gorm:"column:dict_value;not null;comment:字典键值" json:"dictValue"`                             // 字典键值
	DictStatus int32      `gorm:"column:dict_status;not null;comment:字典状态" json:"dictStatus"`                           // 字典状态
	DictSort   int32      `gorm:"column:dict_sort;not null;comment:排序" json:"dictSort"`                                 // 排序
	Remark     string     `gorm:"column:remark;not null;comment:备注信息" json:"remark"`                                    // 备注信息
	IsDefault  int32      `gorm:"column:is_default;not null;comment:是否默认  0：否  1：是" json:"isDefault"`                   // 是否默认  0：否  1：是
	DelFlag    int32      `gorm:"column:del_flag;not null;comment:是否删除  1：已删除  0：正常" json:"delFlag"`                    // 是否删除  1：已删除  0：正常
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建者" json:"createBy"`                                // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by;not null;comment:更新者" json:"updateBy"`                                // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`                                    // 更新时间
}

// TableName SysDictItem's table name
func (*SysDictItem) TableName() string {
	return TableNameSysDictItem
}
