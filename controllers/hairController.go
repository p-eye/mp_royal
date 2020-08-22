package controllers

import (
	"fmt"
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

type Response struct {
	ID uint `json:"hairId"`
	Name string `json:"hairName"`
	ImgUrl string `json:"hairImgUrl"`
	Date map[int][]int
}

type Request struct {
	Year uint
	Month uint
}

func GetHairDate(c *gin.Context) {
	var request Request
	var responses []Response
	c.ShouldBindJSON(&request)

	configs.DB.Table("hair_releases").
		Select("hair_releases.id, name").
		Joins("join monthly_hairs on hair_releases.id = monthly_hairs.hair_release_id").
		Where("monthly_hairs.year = ? and monthly_hairs.month = ?", request.Year, request.Month).
		Find(&responses)

	for i, _ := range responses {
		responses[i].Date = make(map[int][]int)
		rows, err := configs.DB.Table("monthly_hairs").
			Select("year, month").
			Where("hair_release_id = ?", responses[i].ID).
			Rows()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for rows.Next() {
			var year int
			var month int
			rows.Scan(&year, &month)
			responses[i].Date[year] = append(responses[i].Date[year], month)
		}
		/*
		configs.DB.Table("monthly_hairs").
			Select("concat(year, '-', month)").
			Where("hair_release_id = ?", responses[i].ID).
			Scan(&responses[i].Dates)
		 */
		responses[i].ImgUrl = GetHairImg(25017)
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
		"data": responses,
	})
}

func GetHairImg(hairCode int) string {
	return	fmt.Sprintf("https://maplestory.io/api/Character/" +
		"{'itemId':2015,'version':'1105','region':'KMST'}," +
		"{'itemId':12015,'version':'1105','region':'KMST'}," +
		"{'itemId':30000,'version':'1105','region':'KMST'}," +
		"{'itemId':%d,'version':'1105','region':'KMST'}," +
		"{'itemId':1042162,'version':'1105','region':'KMST'}," +
		"{'itemId':1060026,'version':'1105','region':'KMST'}" +
		"/stand1/0?showears=false&showLefEars=false&showHighLefEars=undefined&resize=1&name=&flipX=false&bgColor=0,0,0,0",
		hairCode,
	)
}