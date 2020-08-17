package models

import "gorm.io/gorm"

type MonthlyHair struct {
	gorm.Model
	ReleaseID uint
	Release Release
	Year uint
	Month uint
}
