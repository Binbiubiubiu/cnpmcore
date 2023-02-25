package repository

import (
	"github.com/Binbiubiubiu/cnpmcore/global"
	"github.com/Binbiubiubiu/cnpmcore/repository/model"
)

type PackageRepository struct{}

func (r *PackageRepository) FindPackage(scope string, name string) (*model.Package, error) {
	m := &model.Package{}
	tx := global.DB.Where(&model.Package{Scope: scope, Name: name}).First(m)
	if err := tx.Error; err != nil {
		return nil, err
	}
	// var manifestsDistModel *model.Dist
	// if lo.IsNotEmpty(m.ManifestsDistId) {
	// 	manifestsDistModel := &model.Dist{}
	// 	global.DB.Where(&model.Dist{DistId: m.ManifestsDistId}).First(manifestsDistModel)
	// }
	// var abbreviatedsDistModel *model.Dist
	// if lo.IsNotEmpty(m.AbbreviatedsDistId) {
	// 	abbreviatedsDistModel := &model.Dist{}
	// 	global.DB.Where(&model.Dist{DistId: m.AbbreviatedsDistId}).First(abbreviatedsDistModel)
	// }
	return nil, nil
}
