package model

type PackageDep struct {
	Model
	PackageVersionId string `gorm:"size:24;not null;uniqueIndex:uk_package_version_id_scope_name;comment:package version id;"`
	PackageDepId     string `gorm:"size:24;not null;uniqueIndex:uk_package_dep_id;comment:package dep id;"`
	Scope            string `gorm:"size:214;not null;uniqueIndex:uk_package_version_id_scope_name;comment:package scope;"`
	Name             string `gorm:"type:varchar(214) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL;uniqueIndex:uk_package_version_id_scope_name;comment:package name;"`
	Spec             string `gorm:"size:100;not null;comment:package dep spec;"`
}
