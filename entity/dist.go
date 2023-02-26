package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Dist struct {
	Entity
	DistId    string `json:"dist_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Size      uint32 `json:"size,omitempty"`
	Shasum    string `json:"shasum,omitempty"`
	Integrity string `json:"integrity,omitempty"`
}

func (e *Dist) IntoModel() *model.Dist {
	if e == nil {
		return nil
	}
	var m model.Dist
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.DistId = e.DistId
	m.Name = e.Name
	m.Path = e.Path
	m.Size = e.Size
	m.Shasum = e.Shasum
	m.Integrity = e.Integrity
	return &m
}

func (e *Dist) FromModel(m *model.Dist) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.DistId = m.DistId
	e.Name = m.Name
	e.Path = m.Path
	e.Size = m.Size
	e.Shasum = m.Shasum
	e.Integrity = m.Integrity
}

func CreateDist(e *Dist) *Dist {
	if e == nil {
		e = &Dist{}
	}
	e.DistId = util.CreateObjectId()
	return e
}

func CreateDistFromModel(m *model.Dist) *Dist {
	var e Dist
	e.FromModel(m)
	return &e
}
