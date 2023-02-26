package model

type Dist struct {
	Model
	DistId    string `gorm:"size:24;uniqueIndex:uk_dist_id;not null;comment:dist id;"`
	Name      string `gorm:"size:428;not null;comment:dist name, 2x size of package name;"`
	Path      string `gorm:"type:varchar(767) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL;comment:access path;"`
	Size      uint32 `gorm:"not null;comment:file size;"`
	Shasum    string `gorm:"size:512;not null;comment:dist shasum;"`
	Integrity string `gorm:"size:512;not null;comment:dist integrity;"`
}
