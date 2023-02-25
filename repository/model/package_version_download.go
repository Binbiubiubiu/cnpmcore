package model

type PackageVersionDownload struct {
	Model
	PackageId string `gorm:"size:214;not null;uniqueIndex:uk_year_month_package_id_version;index:idx_packageid_yearmonth;comment:package id, maybe scope name;"`
	Version   string `gorm:"size:256;not null;uniqueIndex:uk_year_month_package_id_version;comment:package version;"`
	YearMonth uint32 `gorm:"not null;uniqueIndex:uk_year_month_package_id_version;index:idx_packageid_yearmonth;index:idx_year_month;comment:YYYYMM format;"`
	D01       uint32 `gorm:"not null;default:0;comment:01 download count;"`
	D02       uint32 `gorm:"not null;default:0;comment:02 download count;"`
	D03       uint32 `gorm:"not null;default:0;comment:03 download count;"`
	D04       uint32 `gorm:"not null;default:0;comment:04 download count;"`
	D05       uint32 `gorm:"not null;default:0;comment:05 download count;"`
	D06       uint32 `gorm:"not null;default:0;comment:06 download count;"`
	D07       uint32 `gorm:"not null;default:0;comment:07 download count;"`
	D08       uint32 `gorm:"not null;default:0;comment:08 download count;"`
	D09       uint32 `gorm:"not null;default:0;comment:09 download count;"`
	D10       uint32 `gorm:"not null;default:0;comment:10 download count;"`
	D11       uint32 `gorm:"not null;default:0;comment:11 download count;"`
	D12       uint32 `gorm:"not null;default:0;comment:12 download count;"`
	D13       uint32 `gorm:"not null;default:0;comment:13 download count;"`
	D14       uint32 `gorm:"not null;default:0;comment:14 download count;"`
	D15       uint32 `gorm:"not null;default:0;comment:15 download count;"`
	D16       uint32 `gorm:"not null;default:0;comment:16 download count;"`
	D17       uint32 `gorm:"not null;default:0;comment:17 download count;"`
	D18       uint32 `gorm:"not null;default:0;comment:18 download count;"`
	D19       uint32 `gorm:"not null;default:0;comment:19 download count;"`
	D20       uint32 `gorm:"not null;default:0;comment:20 download count;"`
	D21       uint32 `gorm:"not null;default:0;comment:21 download count;"`
	D22       uint32 `gorm:"not null;default:0;comment:22 download count;"`
	D23       uint32 `gorm:"not null;default:0;comment:23 download count;"`
	D24       uint32 `gorm:"not null;default:0;comment:24 download count;"`
	D25       uint32 `gorm:"not null;default:0;comment:25 download count;"`
	D26       uint32 `gorm:"not null;default:0;comment:26 download count;"`
	D27       uint32 `gorm:"not null;default:0;comment:27 download count;"`
	D28       uint32 `gorm:"not null;default:0;comment:28 download count;"`
	D29       uint32 `gorm:"not null;default:0;comment:29 download count;"`
	D30       uint32 `gorm:"not null;default:0;comment:30 download count;"`
	D31       uint32 `gorm:"not null;default:0;comment:31 download count;"`
}
