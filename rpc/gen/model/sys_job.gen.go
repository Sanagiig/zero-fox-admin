// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysJob = "sys_job"

// SysJob 职位管理
type SysJob struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                         // 编号
	JobName    string     `gorm:"column:job_name;not null;comment:职位名称" json:"jobName"`                                 // 职位名称
	OrderNum   int32      `gorm:"column:order_num;not null;comment:排序" json:"orderNum"`                                 // 排序
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建人" json:"createBy"`                                // 创建人
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by;not null;comment:更新人" json:"updateBy"`                                // 更新人
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`                                    // 更新时间
	DelFlag    int32      `gorm:"column:del_flag;not null;comment:是否删除  0：已删除  1：正常" json:"delFlag"`                    // 是否删除  0：已删除  1：正常
	Remarks    string     `gorm:"column:remarks;not null;comment:备注" json:"remarks"`                                    // 备注
}

// TableName SysJob's table name
func (*SysJob) TableName() string {
	return TableNameSysJob
}
