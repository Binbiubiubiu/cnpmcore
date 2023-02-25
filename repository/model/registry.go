package model

type Registry struct {
	Model
	RegistryId   string `gorm:"size:24;not null;comment:registry id;"`
	Name         string `gorm:"size:256;uniqueIndex:uk_name;comment:registry name;"`
	Host         string `gorm:"size:4096;comment:registry host;"`
	ChangeStream string `gorm:"size:4096;comment:change stream url;"`
	UserPrefix   string `gorm:"size:256;comment:user prefix;"`
	Type         string `gorm:"size:256;comment:registry type cnpmjsorg/cnpmcore/npm ;"`
}
