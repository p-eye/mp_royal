package models

import "gorm.io/gorm"

type Release struct {
	gorm.Model
	Type string
	Name string
	Year uint
	Month uint
	Sex string
	ImgUrl string
}