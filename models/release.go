package models

import "gorm.io/gorm"

type HairRelease struct {
	gorm.Model
	Name string
	ReleaseYear uint
	ReleaseMonth uint
	Sex string
	MonthlyHairs []MonthlyHair
}

type FaceRelease struct {
	gorm.Model
	Name string
	ReleaseYear uint
	ReleaseMonth uint
	Sex string
	MonthlyFaces []MonthlyFace
}