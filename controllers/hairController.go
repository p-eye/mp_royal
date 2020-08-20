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
	Dates []string
}

type HairDate struct {
	Year uint
	Month uint
	Hairs []Hair
}

type Request struct {
	Year uint
	Month uint
}

func GetHairDate(c *gin.Context) {
	var request Request
	var hairs []Hair
	c.ShouldBindJSON(&request)


	configs.DB.Table("hair_releases").
		Select("hair_releases.id, name").
		Joins("join monthly_hairs on hair_releases.id = monthly_hairs.hair_release_id").
		Where("monthly_hairs.year = ? and monthly_hairs.month = ?", request.Year, request.Month).
		Find(&hairs)

	for i, _ := range hairs {
		configs.DB.Table("monthly_hairs").
			Select("concat(year, '-', month)").
			Where("hair_release_id = ?", hairs[i].ID).
			Scan(&hairs[i].Dates)
	}

/*
	for i, _ := range hairs {
		configs.DB.Table("monthly_hairs").
			Select("year").
			Where("hair_release_id = ?", hairs[i].ID).
			Scan(&hairs[i].Date)
	}
*/
	//configs.DB.Model(&hairs).Association("Comments").Find()
/*
	for i, _ := range hairs {
		configs.DB.Table("monthly_hairs").
			Select("year, month").
			Where("hair_release_id = ?", hairs[i].ID).
			Preload(clause.Associations).
			Scan(&hairs)
	}

*/

	c.JSON(http.StatusOK, gin.H{
		"data": hairs,
	})
}