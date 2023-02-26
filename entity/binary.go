package entity

import (
	"fmt"

	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
)

type Binary struct {
	Entity
	BinaryId               string   `json:"binary_id,omitempty"`
	Category               string   `json:"category,omitempty"`
	Parent                 string   `json:"parent,omitempty"`
	Name                   string   `json:"name,omitempty"`
	IsDir                  bool     `json:"is_dir,omitempty"`
	Size                   uint32   `json:"size,omitempty"`
	Date                   string   `json:"date,omitempty"`
	SourceUrl              string   `json:"source_url,omitempty"`
	IgnoreDownloadStatuses []uint16 `json:"ignore_download_statuses,omitempty"`
}

func (b *Binary) GetStorePath() string {
	return fmt.Sprintf("/binaries/%s%s%s", b.Category, b.Parent, b.Name)
}

func (e *Binary) IntoModel() *model.Binary {
	if e == nil {
		return nil
	}
	var m model.Binary
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.BinaryId = e.BinaryId
	m.Category = e.Category
	m.Parent = e.Parent
	m.Name = e.Name
	m.IsDir = e.IsDir
	m.Size = e.Size
	m.Date = e.Date
	return &m
}

func (e *Binary) FromModel(m *model.Binary) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.BinaryId = m.BinaryId
	e.Category = m.Category
	e.Parent = m.Parent
	e.Name = m.Name
	e.IsDir = m.IsDir
	e.Size = m.Size
	e.Date = m.Date
}

func CreateBinary(e *Binary) *Binary {
	if e == nil {
		e = &Binary{}
	}
	e.BinaryId = util.CreateObjectId()
	return e
}

func CreateBinaryFromModel(m *model.Binary) *Binary {
	var e Binary
	e.FromModel(m)
	return &e
}
