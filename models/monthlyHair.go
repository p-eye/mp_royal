package models

import "gorm.io/gorm"

type MonthlyFace struct {
	gorm.Model
	ReleaseID uint
	Release Release
	Year uint
	Month uint
}
