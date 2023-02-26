package entity

import (
	"time"

	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type PackageVersion struct {
	Entity           `json:"entity,omitempty"`
	PackageId        string     `json:"package_id,omitempty"`
	PackageVersionId string     `json:"package_version_id,omitempty"`
	Version          string     `json:"version,omitempty"`
	AbbreviatedDist  *Dist      `json:"abbreviated_dist,omitempty"`
	ManifestDist     *Dist      `json:"manifest_dist,omitempty"`
	TarDist          *Dist      `json:"tar_dist,omitempty"`
	ReadmeDist       *Dist      `json:"readme_dist,omitempty"`
	PublishTime      *time.Time `json:"publish_time,omitempty"`
}

func (e *PackageVersion) IntoModel() *model.PackageVersion {
	if e == nil {
		return nil
	}
	var m model.PackageVersion
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.PackageId = e.PackageId
	m.PackageVersionId = e.PackageVersionId
	m.Version = e.Version
	// m.AbbreviatedDist = e.AbbreviatedDist
	// m.ManifestDist = e.ManifestDist
	// m.TarDist = e.TarDist
	// m.ReadmeDist = e.ReadmeDist
	m.PublishTime = e.PublishTime
	return &m
}

func (e *PackageVersion) FromModel(m *model.PackageVersion) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.PackageId = m.PackageId
	e.PackageVersionId = m.PackageVersionId
	e.Version = m.Version
	// e.AbbreviatedDist = m.AbbreviatedDist
	// e.ManifestDist = m.ManifestDist
	// e.TarDist = m.TarDist
	// e.ReadmeDist = m.ReadmeDist
	e.PublishTime = m.PublishTime
}

func CreatePackageVersion(e *PackageVersion) *PackageVersion {
	if e == nil {
		e = &PackageVersion{}
	}
	e.PackageVersionId = util.CreateObjectId()
	return e
}

func CreatePackageVersionFromModel(m *model.PackageVersion) *PackageVersion {
	var e PackageVersion
	e.FromModel(m)
	return &e
}
