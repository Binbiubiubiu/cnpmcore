package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/repository/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Scope struct {
	Entity
	Name       string `json:"name,omitempty"`
	ScopeId    string `json:"scope_id,omitempty"`
	RegistryId string `json:"registry_id,omitempty"`
}

func (e *Scope) Into() *model.Scope {
	var m model.Scope
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.Name = e.Name
	m.ScopeId = e.ScopeId
	m.RegistryId = e.RegistryId
	return &m
}

func (e *Scope) From(m *model.Scope) {
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.Name = m.Name
	e.ScopeId = m.ScopeId
	e.RegistryId = m.RegistryId
}

func CreateScope(e *Scope) *Scope {
	e.ScopeId = util.CreateObjectId()
	return e
}

func CreateScopeFromModel(m *model.Scope) *Scope {
	var e Scope
	e.From(m)
	return &e
}
