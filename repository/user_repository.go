package repository

import (
	"github.com/Binbiubiubiu/cnpmcore/core/entity"
	"github.com/Binbiubiubiu/cnpmcore/global"
	"github.com/Binbiubiubiu/cnpmcore/repository/model"
	"github.com/samber/lo"
)

type UserRepository struct{}

func (r *UserRepository) SaveUser(user *entity.User) (err error) {
	var m *model.User
	if lo.IsEmpty(user.ID) {
		defer global.ZLog.Infof("[UserRepository:saveUser:new] id: %s, userId: %s", user.ID, user.UserId)
	} else {
		m = &model.User{}
		tx := global.DB.First(m, user.ID)
		if err = tx.Error; err != nil {
			return
		}
		user.CreateAt = m.CreateAt
	}
	m = user.Into()
	err = global.DB.Save(m).Error
	if err != nil {
		return
	}
	user.From(m)
	return
}

func (r *UserRepository) FindUserByName(name string) (*entity.User, error) {
	m := &model.User{}
	tx := global.DB.Where(&model.User{Name: name}).First(m)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entity.CreateUserFromModel(m), nil
}

func (r *UserRepository) FindUserByUserId(userId string) (*entity.User, error) {
	m := &model.User{}
	tx := global.DB.Where(&model.User{UserId: userId}).First(m)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entity.CreateUserFromModel(m), nil
}

func (r *UserRepository) FindUserAndTokenByTokenKey(tokenKey string) (*entity.Token, *entity.User, error) {
	token, err := r.FindTokenByTokenKey(tokenKey)
	if err != nil {
		return nil, nil, err
	}
	user := &model.User{}
	tx := global.DB.Where(&model.User{UserId: token.UserId}).First(user)
	if err = tx.Error; err != nil {
		return nil, nil, err
	}
	return token, entity.CreateUserFromModel(user), nil
}

func (r *UserRepository) FindTokenByTokenKey(tokenKey string) (*entity.Token, error) {
	m := &model.Token{}
	tx := global.DB.Where(&model.Token{TokenKey: tokenKey}).First(m)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return entity.CreateTokenFromModel(m), nil
}

func (r *UserRepository) SaveToken(token *entity.Token) (err error) {
	var m *model.Token
	if lo.IsEmpty(token.ID) {
		defer global.ZLog.Infof("[UserRepository:saveToken:new] id: %s, tokenId: %s", token.ID, token.TokenId)
	} else {
		m = &model.Token{}
		tx := global.DB.First(m, token.ID)
		if err = tx.Error; err != nil {
			return
		}
		token.CreateAt = m.CreateAt
	}
	m = token.Into()
	err = global.DB.Save(m).Error
	if err != nil {
		return
	}
	token.From(m)
	return
}

func (r *UserRepository) RemoveToken(tokenId string) error {
	tx := global.DB.Delete(&model.Token{TokenId: tokenId})
	defer global.ZLog.Infof("[UserRepository:removeToken:remove] %d rows, tokenId: %s", tx.RowsAffected, tokenId)
	return tx.Error
}

func (r *UserRepository) ListTokens(userId string) ([]*entity.Token, error) {
	var list []*model.Token
	tx := global.DB.Where(&model.User{UserId: userId}).Find(&list)
	if err := tx.Error; err != nil {
		return nil, tx.Error
	}
	return lo.Map(list, func(m *model.Token, i int) *entity.Token {
		return entity.CreateTokenFromModel(m)
	}), nil
}
