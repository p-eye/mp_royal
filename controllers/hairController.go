package controllers

import (
	"github.com/gin-gonic/gin"
	"mp_royal/configs"
	"net/http"
)

type MonthlyHair struct {
	HairReleaseID uint
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