package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type PackageVersionBlock struct {
	Entity                `json:"entity,omitempty"`
	PackageId             string `json:"package_id,omitempty"`
	PackageVersionBlockId string `json:"package_version_block_id,omitempty"`
	Reason                string `json:"reason,omitempty"`
	Version               string `json:"version,omitempty"`
}

func (e *PackageVersionBlock) IntoModel() *model.PackageVersionBlock {
	var m model.PackageVersionBlock
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.PackageId = e.PackageId
	m.PackageVersionBlockId = e.PackageVersionBlockId
	m.Reason = e.Reason
	m.Version = e.Version
	return &m
}

func (e *PackageVersionBlock) FromModel(m *model.PackageVersionBlock) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.PackageId = m.PackageId
	e.PackageVersionBlockId = m.PackageVersionBlockId
	e.Reason = m.Reason
	e.Version = m.Version
}

func CreatePackageVersionBlock(e *PackageVersionBlock) *PackageVersionBlock {
	if e == nil {
		e = &PackageVersionBlock{}
	}
	e.PackageVersionBlockId = util.CreateObjectId()
	return e
}

func CreatePackageVersionBlockFromModel(m *model.PackageVersionBlock) *PackageVersionBlock {
	var e PackageVersionBlock
	e.FromModel(m)
	return &e
}
