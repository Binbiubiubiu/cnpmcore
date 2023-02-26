package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type PackageVersionManifest struct {
	Entity
	PackageId                string `json:"package_id,omitempty"`
	PackageVersionId         string `json:"package_version_id,omitempty"`
	PackageVersionManifestId string `json:"package_version_manifest_id,omitempty"`
	Manifest                 any    `json:"manifest,omitempty"`
}

func (e *PackageVersionManifest) IntoModel() *model.PackageVersionManifest {
	if e == nil {
		return nil
	}
	var m model.PackageVersionManifest
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.PackageId = e.PackageId
	m.PackageVersionId = e.PackageVersionId
	m.PackageVersionManifestId = e.PackageVersionManifestId
	m.Manifest = e.Manifest
	return &m
}

func (e *PackageVersionManifest) FromModel(m *model.PackageVersionManifest) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.PackageId = m.PackageId
	e.PackageVersionId = m.PackageVersionId
	e.PackageVersionManifestId = m.PackageVersionManifestId
	e.Manifest = m.Manifest
}

func CreatePackageVersionManifest(e *PackageVersionManifest) *PackageVersionManifest {
	if e == nil {
		e = &PackageVersionManifest{}
	}
	e.PackageVersionManifestId = util.CreateObjectId()
	return e
}

func CreatePackageVersionManifestFromModel(m *model.PackageVersionManifest) *PackageVersionManifest {
	var e PackageVersionManifest
	e.FromModel(m)
	return &e
}
