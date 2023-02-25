package model

type PackageVersionManifest struct {
	Model
	PackageId                string `gorm:"size:24;not null;index:idx_package_id;comment:package id;"`
	PackageVersionId         string `gorm:"size:24;not null;uniqueIndex:uk_package_version_id;comment:package version id;"`
	PackageVersionManifestId string `gorm:"size:24;not null;uniqueIndex:uk_package_version_manifest_id;comment:package version manifest id;"`
	Manifest                 any    `gorm:"serializer:json;type:json NOT NULL;comment:manifest JSON, including README text;"`
}
