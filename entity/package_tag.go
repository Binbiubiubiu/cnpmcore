package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type PackageTag struct {
	Entity
	PackageId    string `json:"package_id,omitempty"`
	PackageTagId string `json:"package_tag_id,omitempty"`
	Tag          string `json:"tag,omitempty"`
	Version      string `json:"version,omitempty"`
}

func (e *PackageTag) IntoModel() *model.PackageTag {
	if e == nil {
		return nil
	}
	var m model.PackageTag
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.PackageId = e.PackageId
	m.PackageTagId = e.PackageTagId
	m.Tag = e.Tag
	m.Version = e.Version
	return &m
}

func (e *PackageTag) FromModel(m *model.PackageTag) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.PackageId = m.PackageId
	e.PackageTagId = m.PackageTagId
	e.Tag = m.Tag
	e.Version = m.Version
}

func CreatePackageTag(e *PackageTag) *PackageTag {
	if e == nil {
		e = &PackageTag{}
	}
	e.PackageTagId = util.CreateObjectId()
	return e
}

func CreatePackageTagFromModel(m *model.PackageTag) *PackageTag {
	var e PackageTag
	e.FromModel(m)
	return &e
}
