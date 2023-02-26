package model

type Maintainer struct {
	Model
	PackageId string `gorm:"size:24;not null;uniqueIndex:uk_package_id_user_id;index:idx_package_id;comment:package id;"`
	UserId    string `gorm:"size:24;not null;uniqueIndex:uk_package_id_user_id;index:idx_user_id;comment:user id;"`
}
