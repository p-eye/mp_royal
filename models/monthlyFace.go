package models

import "gorm.io/gorm"

type MonthlyHair struct {
	gorm.Model
	Release Release
	Year uint
	Month uint
}
