package model

type Scope struct {
	Model
	Name       string `gorm:"size:214;uniqueIndex:uk_name;comment:scope id;"`
	ScopeId    string `gorm:"size:24;not null;comment:scope name;"`
	RegistryId string `gorm:"size:24;not null;comment:registry id;"`
}
