package model

type PackageTag struct {
	Model
	PackageId    string `gorm:"size:24;not null;uniqueIndex:uk_package_tag;comment:package id;"`
	PackageTagId string `gorm:"size:24;not null;uniqueIndex:uk_package_tag_id;comment:package tag id;"`
	Tag          string `gorm:"size:214;not null;uniqueIndex:uk_package_tag;comment:package tag;"`
	Version      string `gorm:"size:256;not null;comment:package version;"`
}
