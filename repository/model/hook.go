package model

import (
	"time"
)

type Hook struct {
	ID           uint      `gorm:"index:idx_type_name_id;comment:primary key"`
	CreateAt     time.Time `gorm:"column:gmt_create;not null;comment:create time"`
	UpdateAt     time.Time `gorm:"column:gmt_modified;not null;comment:modified time"`
	HookId       string    `gorm:"size:24;not null;comment:hook id;"`
	Type         string    `gorm:"size:20;uniqueIndex:uk_type_name_owner_id;index:idx_type_name_id;not null;comment:hook type, scope, name, owner;"`
	Name         string    `gorm:"type:varchar(428) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;uniqueIndex:uk_type_name_owner_id;index:idx_type_name_id;comment:hook name;"`
	OwnerId      string    `gorm:"size:24;uniqueIndex:uk_type_name_owner_id;not null;comment:hook owner id;"`
	Endpoint     string    `gorm:"size:2048;not null;comment:hook url;"`
	Secret       string    `gorm:"size:200;not null;comment:sign secret;"`
	LatestTaskId string    `gorm:"size:24;comment:latest task id;"`
	Enable       bool      `gorm:"not null;default:0;comment:hook is enable not, 1: true, other: false;"`
}
