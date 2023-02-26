package entity

import (
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type User struct {
	Entity
	UserId            string   `json:"user_id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Email             string   `json:"email,omitempty"`
	PasswordSalt      string   `json:"password_salt,omitempty"`
	PasswordIntegrity string   `json:"password_integrity,omitempty"`
	Ip                string   `json:"ip,omitempty"`
	IsPrivate         bool     `json:"is_private,omitempty"`
	Scopes            []string `json:"scopes,omitempty"`
}

func (e *User) IntoModel() *model.User {
	if e == nil {
		return nil
	}
	var m model.User
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.UserId = e.UserId
	m.Name = e.Name
	m.Email = e.Email
	m.PasswordSalt = e.PasswordSalt
	m.PasswordIntegrity = e.PasswordIntegrity
	m.Ip = e.Ip
	m.IsPrivate = e.IsPrivate
	m.Scopes = e.Scopes
	return &m
}

func (e *User) FromModel(m *model.User) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.UserId = m.UserId
	e.Name = m.Name
	e.Email = m.Email
	e.PasswordSalt = m.PasswordSalt
	e.PasswordIntegrity = m.PasswordIntegrity
	e.Ip = m.Ip
	e.IsPrivate = m.IsPrivate
	e.Scopes = m.Scopes
}

func CreateUser(e *User) *User {
	if e == nil {
		e = &User{}
	}
	e.UserId = util.CreateObjectId()
	return e
}

func CreateUserFromModel(m *model.User) *User {
	var e User
	e.FromModel(m)
	return &e
}
