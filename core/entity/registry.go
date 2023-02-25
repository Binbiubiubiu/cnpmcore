package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/repository/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Registry struct {
	Entity
	Name         string `json:"name,omitempty"`
	RegistryId   string `json:"registry_id,omitempty"`
	Host         string `json:"host,omitempty"`
	ChangeStream string `json:"change_stream,omitempty"`
	UserPrefix   string `json:"user_prefix,omitempty"`
	Type         string `json:"type,omitempty"`
}

func (e *Registry) Into() *model.Registry {
	var m model.Registry
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.Name = e.Name
	m.RegistryId = e.RegistryId
	m.Host = e.Host
	m.ChangeStream = e.ChangeStream
	m.UserPrefix = e.UserPrefix
	m.Type = e.Type
	return &m
}

func (e *Registry) From(m *model.Registry) {
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.Name = m.Name
	e.RegistryId = m.RegistryId
	e.Host = m.Host
	e.ChangeStream = m.ChangeStream
	e.UserPrefix = m.UserPrefix
	e.Type = m.Type
}

func CreateRegistry(e *Registry) *Registry {
	e.RegistryId = util.CreateObjectId()
	return e
}

func CreateRegistryFromModel(m *model.Registry) *Registry {
	var e Registry
	e.From(m)
	return &e
}
