package entity

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Hook struct {
	Entity
	HookId       string `json:"hook_id,omitempty"`
	Type         string `json:"type,omitempty"`
	OwnerId      string `json:"owner_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Endpoint     string `json:"endpoint,omitempty"`
	Secret       string `json:"secret,omitempty"`
	Enable       bool   `json:"enable,omitempty"`
	LatestTaskId string `json:"latest_task_id,omitempty"`
}

func (e *Hook) SignPayload(payload any) (payloadStr string, digest string, err error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return
	}
	payloadStr = string(b)
	h := hmac.New(sha256.New, []byte(e.Secret))
	digest = hex.EncodeToString(h.Sum(b))
	return
}

func (e *Hook) IntoModel() *model.Hook {
	if e == nil {
		return nil
	}
	var m model.Hook
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.HookId = e.HookId
	m.Type = e.Type
	m.OwnerId = e.OwnerId
	m.Name = e.Name
	m.Endpoint = e.Endpoint
	m.Secret = e.Secret
	m.Enable = e.Enable
	m.LatestTaskId = e.LatestTaskId
	return &m
}

func (e *Hook) FromModel(m *model.Hook) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.HookId = m.HookId
	e.Type = m.Type
	e.OwnerId = m.OwnerId
	e.Name = m.Name
	e.Endpoint = m.Endpoint
	e.Secret = m.Secret
	e.Enable = m.Enable
	e.LatestTaskId = m.LatestTaskId
}

func CreateHook(e *Hook) *Hook {
	if e == nil {
		e = &Hook{}
	}
	e.HookId = util.CreateObjectId()
	e.Enable = true
	return e
}

func CreateHookFromModel(m *model.Hook) *Hook {
	var e Hook
	e.FromModel(m)
	return &e
}
