package model

type PackageVersionBlock struct {
	Model
	PackageId             string `gorm:"size:24;not null;uniqueIndex:uk_name_version;comment:package id;"`
	PackageVersionBlockId string `gorm:"size:24;not null;uniqueIndex:uk_package_version_block_id;comment:package version block id;"`
	Version               string `gorm:"size:256;not null;uniqueIndex:uk_name_version;comment:package version, \"*\" meaning all versions;"`
	Reason                string `gorm:"type:longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL;comment:block reason;"`
}
