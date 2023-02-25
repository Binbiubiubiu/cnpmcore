package model

import (
	"time"
)

type Task struct {
	ID               uint           `gorm:"comment:primary key"`
	CreateAt         time.Time      `gorm:"column:gmt_create;not null;comment:create time"`
	UpdateAt         time.Time      `gorm:"column:gmt_modified;not null;index:idx_type_state_gmt_modified;index:idx_gmt_modified;comment:modified time"`
	TaskId           string         `gorm:"size:24;not null;uniqueIndex:uk_task_id;comment:task id;"`
	Type             string         `gorm:"size:20;not null;index:idx_type_state_target_name;index:idx_type_state_gmt_modified;comment:scope id;"`
	State            string         `gorm:"size:20;not null;index:idx_type_state_target_name;index:idx_type_state_gmt_modified;comment:task state;"`
	TargetName       string         `gorm:"type:varchar(214) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;index:idx_type_state_target_name;comment:target name, like package name / user name;"`
	AuthorId         string         `gorm:"size:24;not null;comment:create task user id;"`
	AuthorIp         string         `gorm:"size:100;not null;comment:create task user request ip;"`
	Data             map[string]any `gorm:"serializer:json;type:json NULL;comment:task params;"`
	LogPath          string         `gorm:"size:512;not null;comment:access path;"`
	LogStorePosition string         `gorm:"size:10;not null;comment:cloud store disk position;"`
	Attempts         uint32         `gorm:"default:0;comment:task execute attempts times;"`
	Error            string         `gorm:"type:longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL;comment:error description;"`
	BizId            string         `gorm:"size:100;uniqueIndex:uk_biz_id;comment:unique biz id to keep task unique;"`
}
