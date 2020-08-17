package models

import "gorm.io/gorm"

type Release struct {
	gorm.Model
	Type string
	Name string
	ReleaseYear uint
	ReleaseMonth uint
	Sex string
	ImgUrl string
}