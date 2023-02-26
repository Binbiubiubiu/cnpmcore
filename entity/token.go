package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Token struct {
	Entity
	TokenId       string   `json:"token_id,omitempty"`
	TokenMark     string   `json:"token_mark,omitempty"`
	TokenKey      string   `json:"token_key,omitempty"`
	IsReadonly    bool     `json:"is_readonly,omitempty"`
	IsAutomation  bool     `json:"is_automation,omitempty"`
	UserId        string   `json:"user_id,omitempty"`
	CidrWhitelist []string `json:"cidr_whitelist,omitempty"`
}

func (e *Token) IntoModel() *model.Token {
	if e == nil {
		return nil
	}
	var m model.Token
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.TokenId = e.TokenId
	m.TokenMark = e.TokenMark
	m.TokenKey = e.TokenKey
	m.IsReadonly = e.IsReadonly
	m.IsAutomation = e.IsAutomation
	m.UserId = e.UserId
	m.CidrWhitelist = e.CidrWhitelist
	return &m
}

func (e *Token) FromModel(m *model.Token) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.TokenId = m.TokenId
	e.TokenMark = m.TokenMark
	e.TokenKey = m.TokenKey
	e.IsReadonly = m.IsReadonly
	e.IsAutomation = m.IsAutomation
	e.UserId = m.UserId
	e.CidrWhitelist = m.CidrWhitelist
}

func CreateToken(e *Token) *Token {
	if e == nil {
		e = &Token{}
	}
	e.UserId = util.CreateObjectId()
	return e
}

func CreateTokenFromModel(m *model.Token) *Token {
	var e Token
	e.FromModel(m)
	return &e
}
