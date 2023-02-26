package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Scope struct {
	Entity
	Name       string `json:"name,omitempty"`
	ScopeId    string `json:"scope_id,omitempty"`
	RegistryId string `json:"registry_id,omitempty"`
}

func (e *Scope) IntoModel() *model.Scope {
	if e == nil {
		return nil
	}
	var m model.Scope
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.Name = e.Name
	m.ScopeId = e.ScopeId
	m.RegistryId = e.RegistryId
	return &m
}

func (e *Scope) FromModel(m *model.Scope) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.Name = m.Name
	e.ScopeId = m.ScopeId
	e.RegistryId = m.RegistryId
}

func CreateScope(e *Scope) *Scope {
	if e == nil {
		e = &Scope{}
	}
	e.ScopeId = util.CreateObjectId()
	return e
}

func CreateScopeFromModel(m *model.Scope) *Scope {
	var e Scope
	e.FromModel(m)
	return &e
}
