package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID       uint      `gorm:"comment:primary key"`
	CreateAt time.Time `gorm:"column:gmt_create;not null;comment:create time"`
	UpdateAt time.Time `gorm:"column:gmt_modified;not null;comment:modified time"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = time.Now()
	m.UpdateAt = time.Now()
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.Omit("id", "gmt_create")
	m.UpdateAt = time.Now()
	return
}
