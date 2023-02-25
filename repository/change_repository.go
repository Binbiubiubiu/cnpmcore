package repository

import (
	"github.com/Binbiubiubiu/cnpmcore/core/entity"
	"github.com/Binbiubiubiu/cnpmcore/global"
	"github.com/Binbiubiubiu/cnpmcore/repository/model"
	"github.com/samber/lo"
)

type ChangeRepository struct{}

func (r *ChangeRepository) AddChange(change *entity.Change) (err error) {
	m := change.Into()
	err = global.DB.Create(m).Error
	if err != nil {
		return
	}
	change.From(m)
	return
}

func (r *ChangeRepository) Query(since uint, limit int) ([]*entity.Change, error) {
	var list []*model.Change
	tx := global.DB.Where("id >= ?", since).Order("id").Limit(limit).Find(&list)
	return lo.Map(list, func(m *model.Change, i int) *entity.Change {
		return entity.CreateChangeFromModel(m)
	}), tx.Error
}

func (r *ChangeRepository) GetLastChange() (*entity.Change, error) {
	c := &model.Change{}
	tx := global.DB.Order("id desc").Limit(1).Find(c)
	return entity.CreateChangeFromModel(c), tx.Error
}
