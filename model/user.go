package model

type User struct {
	Model
	UserId            string   `gorm:"size:24;not null;uniqueIndex:uk_user_id;comment:user id;"`
	Name              string   `gorm:"size:100;not null;uniqueIndex:uk_name;comment:user name;"`
	Email             string   `gorm:"size:400;not null;comment:user email;"`
	PasswordSalt      string   `gorm:"size:100;not null;comment:password salt;"`
	PasswordIntegrity string   `gorm:"size:512;not null;comment:password integrity;"`
	Ip                string   `gorm:"size:100;not null;comment:user login request ip;"`
	IsPrivate         bool     `gorm:"type:tinyint NOT NULL DEFAULT 1;comment:private user or not, 1: true, other: false;"`
	Scopes            []string `gorm:"type:json NULL;comment:white scope list, [\"@cnpm\", \"@foo\"];"`
}
