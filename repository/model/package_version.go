package model

import (
	"time"
)

type PackageVersion struct {
	Model
	PackageId         string    `gorm:"size:24;not null;uniqueIndex:uk_package_id_version;comment:package id;"`
	PackageVersionId  string    `gorm:"size:24;not null;uniqueIndex:uk_package_version_id;comment:package version id;"`
	Version           string    `gorm:"size:256;not null;uniqueIndex:uk_package_id_version;comment:package version;"`
	AbbreviatedDistId string    `gorm:"size:24;not null;comment:abbreviated manifest dist id;"`
	ManifestDistId    string    `gorm:"size:24;not null;comment:manifest dist id;"`
	TarDistId         string    `gorm:"size:24;not null;comment:package id;"`
	ReadmeDistId      string    `gorm:"size:24;not null;comment:readme dist id;"`
	PublishTime       time.Time `gorm:"not null;comment:publish time;"`
}
