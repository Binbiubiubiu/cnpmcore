package model

type Binary struct {
	Model
	BinaryId string `gorm:"size:24;not null;uniqueIndex:uk_binary_id;comment:binary id;"`
	Category string `gorm:"type:varchar(50) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT 'binary category, e.g.: node, sass';uniqueIndex:uk_category_parent_name;index:idx_category_parent;"`
	Parent   string `gorm:"type:varchar(500) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT 'binary parent name, e.g.: /, /v1.0.0/, /v1.0.0/docs/';uniqueIndex:uk_category_parent_name;index:idx_category_parent;"`
	Name     string `gorm:"type:varchar(200) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT 'binary name, dir should ends with /';uniqueIndex:uk_category_parent_name;"`
	IsDir    bool   `gorm:"not null;default:0;comment:is dir or not, 1: true, other: false;"`
	Size     uint32 `gorm:"not null;comment:file size;"`
	Date     string `gorm:"size:100;not null;comment:date display string;"`
}
