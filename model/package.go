package model

type Package struct {
	Model
	PackageId          string `gorm:"size:24;not null;uniqueIndex:uk_package_id;comment:package id;"`
	RegistryId         string `gorm:"size:24;not null;comment:source registry;"`
	Scope              string `gorm:"size:214;not null;uniqueIndex:uk_scope_name;comment:package name, empty string meaning no scope;"`
	Name               string `gorm:"type:varchar(214) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;uniqueIndex:uk_scope_name;comment:package name;"`
	IsPrivate          bool   `gorm:"not null;default:0;comment:private pkg or not, 1: true, other: false;"`
	Description        string `gorm:"type:varchar(10240) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL;comment:package description;"`
	AbbreviatedsDistId string `gorm:"size:24;comment:all abbreviated manifests dist id;"`
	ManifestsDistId    string `gorm:"size:24;comment:all full manifests dist id;"`
	AbbreviatedsDist   *Dist  `gorm:"-"`
	ManifestsDist      *Dist  `gorm:"-"`
}
