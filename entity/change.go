package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Change struct {
	Entity
	ChangeId   string `json:"change_id,omitempty"`
	Type       string `json:"type,omitempty"`
	TargetName string `json:"target_name,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func (e *Change) IntoModel() *model.Change {
	if e == nil {
		return nil
	}
	var m model.Change
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.ChangeId = e.ChangeId
	m.Type = e.Type
	m.TargetName = e.TargetName
	m.Data = e.Data
	return &m
}

func (e *Change) FromModel(m *model.Change) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.ChangeId = m.ChangeId
	e.Type = m.Type
	e.TargetName = m.TargetName
	e.Data = m.Data
}

func CreateChange(e *Change) *Change {
	if e == nil {
		e = &Change{}
	}
	e.ChangeId = util.CreateObjectId()
	return e
}

func CreateChangeFromModel(m *model.Change) *Change {
	var e Change
	e.FromModel(m)
	return &e
}
