package models

import "gorm.io/gorm"

type MonthlyFace struct {
	gorm.Model
	Release Release
	Year uint
	Month uint
}
