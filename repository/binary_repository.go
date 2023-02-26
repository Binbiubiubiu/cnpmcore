package repository

import (
	"github.com/Binbiubiubiu/cnpmcore/entity"
	"github.com/Binbiubiubiu/cnpmcore/global"
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/samber/lo"
)

type BinaryRepository struct{}

func (r *BinaryRepository) SaveBinary(binary *entity.Binary) (err error) {
	var m *model.Binary
	if lo.IsEmpty(binary.ID) {
		defer global.ZLog.Infof("[BinaryRepository:saveBinary:new] id: %v, binaryId: %s", binary.ID, binary.BinaryId)
	} else {
		m = &model.Binary{}
		tx := global.DB.First(m, binary.ID)
		if err = tx.Error; err != nil {
			return
		}
		binary.CreateAt = m.CreateAt
	}
	m = binary.IntoModel()
	err = global.DB.Save(m).Error
	if err != nil {
		return
	}
	binary.FromModel(m)
	return
}

func (r *BinaryRepository) FindBinary(category string, parent string, name string) (*entity.Binary, error) {
	m := &model.Binary{}
	tx := global.DB.Where(&model.Binary{Category: category, Parent: parent, Name: name}).First(m)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return entity.CreateBinaryFromModel(m), nil
}

func (r *BinaryRepository) ListBinaries(category string, parent string) ([]*entity.Binary, error) {
	var list []*model.Binary
	tx := global.DB.Where(&model.Binary{Category: category, Parent: parent}).Find(&list)
	if err := tx.Error; err != nil {
		return nil, tx.Error
	}
	return lo.Map(list, func(m *model.Binary, i int) *entity.Binary {
		return entity.CreateBinaryFromModel(m)
	}), nil
}
