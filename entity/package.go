package entity

import (
	"fmt"

	"github.com/Binbiubiubiu/cnpmcore/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
	"github.com/samber/lo"
)

const (
	distNameAbbreviated          = "abbreviated.json"
	distNameManifest             = "package.json"
	distNameReadme               = "readme.md"
	distNameFullManifests        = "full_manifests.json"
	distNameAbbreviatedManifests = "abbreviated_manifests.json"
)

type Package struct {
	Entity
	Scope            string `json:"scope,omitempty"`
	Name             string `json:"name,omitempty"`
	PackageId        string `json:"package_id,omitempty"`
	IsPrivate        bool   `json:"is_private,omitempty"`
	Description      string `json:"description,omitempty"`
	AbbreviatedsDist *Dist  `json:"abbreviateds_dist"`
	ManifestsDist    *Dist  `json:"manifests_dist"`
	RegistryId       string `json:"registry_id,omitempty"`
}

func (e *Package) FullName() string {
	return util.GetFullname(e.Scope, e.Name)
}

func (e *Package) CreateAbbreviated(version string, info *FileInfo) *Dist {
	return e.createDist(distNameAbbreviated, info, version)
}

func (e *Package) CreateManifest(version string, info *FileInfo) *Dist {
	return e.createDist(distNameManifest, info, version)
}

func (e *Package) CreateReadme(version string, info *FileInfo) *Dist {
	return e.createDist(distNameReadme, info, version)
}

func (e *Package) CreateTar(version string, info *FileInfo) *Dist {
	return e.createDist(fmt.Sprintf("%s-%s.tgz", e.Name, version), info, version)
}

func (e *Package) CreateFullMainfests(info *FileInfo) *Dist {
	return e.createDist(distNameFullManifests, info)
}

func (e *Package) CreateAbbreviatedManifests(info *FileInfo) *Dist {
	return e.createDist(distNameAbbreviatedManifests, info)
}

func (e *Package) distDir(filename string, args ...string) string {
	var version string
	if len(args) > 0 && lo.IsNotEmpty(args[0]) {
		version = "/" + args[0]
	}
	return fmt.Sprintf("/packages/%s%s/%s", e.FullName(), version, filename)
}

func (e *Package) createDist(name string, info *FileInfo, args ...string) *Dist {
	return CreateDist(&Dist{
		Name:      name,
		Size:      info.Size,
		Shasum:    info.Shasum,
		Integrity: info.Integrity,
		Path:      e.distDir(name, args...),
	})
}

func (e *Package) IntoModel() *model.Package {
	if e == nil {
		return nil
	}
	var m model.Package
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.Scope = e.Scope
	m.Name = e.Name
	m.PackageId = e.PackageId
	m.IsPrivate = e.IsPrivate
	m.Description = e.Description
	m.AbbreviatedsDist = e.AbbreviatedsDist.IntoModel()
	m.ManifestsDist = e.ManifestsDist.IntoModel()
	m.RegistryId = e.RegistryId
	return &m
}

func (e *Package) FromModel(m *model.Package) {
	if e == nil || m == nil {
		return
	}
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.Scope = m.Scope
	e.Name = m.Name
	e.PackageId = m.PackageId
	e.IsPrivate = m.IsPrivate
	e.Description = m.Description
	e.AbbreviatedsDist = CreateDistFromModel(m.AbbreviatedsDist)
	e.ManifestsDist = CreateDistFromModel(m.ManifestsDist)
	e.RegistryId = m.RegistryId
}

func CreatePackage(e *Package) *Package {
	if e == nil {
		e = &Package{}
	}
	e.PackageId = util.CreateObjectId()
	return e
}

func CreatePackageFromModel(m *model.Package) *Package {
	var e Package
	e.FromModel(m)
	return &e
}

type FileInfo struct {
	Size      uint32
	Shasum    string
	Integrity string
}
