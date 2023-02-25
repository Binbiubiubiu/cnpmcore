package model

type Token struct {
	Model
	TokenId       string   `gorm:"size:24;not null;uniqueIndex:uk_token_id;comment:token id;"`
	TokenMark     string   `gorm:"size:20;not null;comment:token mark value;"`
	TokenKey      string   `gorm:"size:200;not null;uniqueIndex:uk_token_key;comment:token value sha512 hex;"`
	IsReadonly    bool     `gorm:"not null;default:0;comment:readonly token or not, 1: true, other: false;"`
	IsAutomation  bool     `gorm:"not null;default:0;comment:automation token or not, 1: true, other: false;"`
	UserId        string   `gorm:"size:24;not null;index:idx_user_id;comment:user idd;"`
	CidrWhitelist []string `gorm:"serializer:json;type:json NOT NULL;comment:ip list, [\"127.0.0.1\"];"`
}
