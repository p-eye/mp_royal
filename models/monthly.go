package models

import "gorm.io/gorm"

type MonthlyHair struct {
	gorm.Model
	HairReleaseID uint
	Year uint
	Month uint
}

type MonthlyFace struct {
	gorm.Model
	FaceReleaseID uint
	Year uint
	Month uint
}
