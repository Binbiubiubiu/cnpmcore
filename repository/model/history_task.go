package model

type HistoryTask struct {
	Model
	TaskId           string         `gorm:"size:24;uniqueIndex:uk_task_id;not null;comment:task id;"`
	Type             string         `gorm:"size:20;not null;comment:task type;"`
	State            string         `gorm:"size:20;not null;comment:task state;"`
	TargetName       string         `gorm:"size:214;not null;comment:target name, like package name / user name;"`
	AuthorId         string         `gorm:"size:24;not null;comment:create task user id;"`
	AuthorIp         string         `gorm:"size:100;not null;comment:create task user request ip;"`
	Data             map[string]any `gorm:"serializer:json;type:json NULL;comment:task params;"`
	LogPath          string         `gorm:"size:512;not null;comment:access path;"`
	LogStorePosition string         `gorm:"size:10;not null;comment:cloud store disk position;"`
	Attempts         uint32         `gorm:"default:0;comment:task execute attempts times;"`
	Error            string         `gorm:"type:longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL;comment:error description;"`
}
