package controllers

import (
	"github.com/gin-gonic/gin"
	"mp_royal/configs"
	"net/http"
)

type MonthlyHair struct {
	HairReleaseID uint `json:"hairId"`
	Year uint `json:"year"`
	Month uint `json:"month"`
}

type HairRelease struct {
	ID              uint   `json:"hairId"`
	Name            string `json:"hairName"`
	ImgUrl          string `json:"hairImgUrl"`
	MonthlyHairs []MonthlyHair
}

func GetHairCycle(c *gin.Context) {
	var hairRelease HairRelease
	var monthlyHairs []MonthlyHair
	id := c.PostForm("id")

	configs.DB.Where("id = ?", id).First(&hairRelease)
	configs.DB.Where("hair_release_id = ?", id).Find(&monthlyHairs)
	hairRelease.MonthlyHairs = monthlyHairs

	//configs.DB.Model(&hairRelease).Association("monthly_hairs").Find(&hairRelease.MonthlyHairs)
	c.JSON(http.StatusOK, gin.H{
		"data": hairRelease,
	})
}

type Hair struct {
	ID uint `json:"hairId"`
	Name string `json:"hairName"`
	ImgUrl string `json:"hairImageURl"`
}

type HairDate struct {
	Year uint
	Month uint
	Hairs []Hair
}


func GetHairDate(c *gin.Context) {
	var hairDate HairDate

	c.ShouldBindJSON(&hairDate)
	configs.DB.Table("hair_releases").
		Select("hair_releases.id, name, img_url").
		Joins("join monthly_hairs on hair_releases.id = monthly_hairs.hair_release_id").
		Where("monthly_hairs.year = ? and monthly_hairs.month = ?", hairDate.Year, hairDate.Month).
		Scan(&hairDate.Hairs)
	c.JSON(http.StatusOK, gin.H{
		"data": hairDate,
	})
}