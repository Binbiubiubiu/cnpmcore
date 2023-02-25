package model

type Change struct {
	Model
	ChangeId   string `gorm:"size:24;uniqueIndex:uk_change_id;not null;comment:change id;"`
	Type       string `gorm:"size:50;not null;comment:change type;"`
	TargetName string `gorm:"size:214;not null;comment:target name, like package name / user name;"`
	Data       any    `gorm:"serializer:json;type:json NULL;comment:change params;"`
}
