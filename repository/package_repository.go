package repository

import (
	"github.com/Binbiubiubiu/cnpmcore/entity"
	"github.com/Binbiubiubiu/cnpmcore/global"
	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/samber/lo"
)

type PackageRepository struct{}

func (r *PackageRepository) FindPackage(scope string, name string) (*entity.Package, error) {
	m := &model.Package{}
	err := global.DB.Where(&model.Package{Scope: scope, Name: name}).First(m).Error
	if err != nil {
		return nil, err
	}
	if lo.IsNotEmpty(m.ManifestsDistId) {
		m.ManifestsDist = &model.Dist{}
		global.DB.Where(&model.Dist{DistId: m.ManifestsDistId}).First(m.ManifestsDist)
	}
	if lo.IsNotEmpty(m.AbbreviatedsDistId) {
		m.AbbreviatedsDist = &model.Dist{}
		global.DB.Where(&model.Dist{DistId: m.AbbreviatedsDistId}).First(m.AbbreviatedsDist)
	}
	e := entity.CreatePackageFromModel(m)
	if lo.IsEmpty(e.AbbreviatedsDist.ID) {
		e.AbbreviatedsDist = nil
	}
	if lo.IsEmpty(e.ManifestsDist.ID) {
		e.ManifestsDist = nil
	}
	return e, nil
}

func (r *PackageRepository) FindPackageId(scope string, name string) (string, error) {
	m := &model.Package{}
	err := global.DB.Where(&model.Package{Scope: scope, Name: name}).First(m).Error
	if err != nil {
		return "", err
	}
	return m.PackageId, nil
}

func (r *PackageRepository) SavePackage(pkg *entity.Package) (err error) {
	var m *model.Package
	if lo.IsEmpty(pkg.ID) {
		defer global.ZLog.Infof("[PackageRepository:savePackage:new] id: %v, PackageId: %s", pkg.ID, pkg.PackageId)
	} else {
		m = &model.Package{}
		tx := global.DB.First(m, pkg.ID)
		if err = tx.Error; err != nil {
			return
		}
		pkg.CreateAt = m.CreateAt
	}
	m = pkg.IntoModel()
	err = global.DB.Save(m).Error
	if err != nil {
		return
	}
	pkg.FromModel(m)
	return
}
