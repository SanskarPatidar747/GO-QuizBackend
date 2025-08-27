package quizFunction

import (
	"fmt"
	//"github.com/gin-contrib/cors"                        // Why do we need this package?
	"github.com/gin-gonic/gin" // Using gin as microframework
	"github.com/jinzhu/gorm"   //Using gorm as orm
	_ "modernc.org/sqlite"     //Using modernc.org/sqlite as db driver
)

type Quiz struct {
	ID       uint   `json:"id"`
	Title    string `gorm:"type:varchar(100)" json:"title"`
	Genre_id uint   `json:"genre_id"`
}

func GetAllQuizs(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	genre_id := c.Params.ByName("genre_id")
	var quiz []Quiz
	if check := db.Where("genre_id = ?", genre_id).Find(&quiz).Error; check != nil {
		c.Header("access-control-allow-origin", "*")
		c.AbortWithStatus(404)
		fmt.Println(check)
	} else {
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, quiz)
	}
}

func GetQuiz(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	id := c.Params.ByName("id")
	var quiz []Quiz
	if check := db.Where("id = ?", id).First(&quiz).Error; check != nil {
		c.AbortWithStatus(404)
		fmt.Println(check)
	} else {
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, quiz)
	}
}

func AddQuiz(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()

	var quiz Quiz
	if err = c.BindJSON(&quiz); err != nil {
		c.Header("access-control-allow-origin", "*")
		c.JSON(400, err)
		return
	}
	fmt.Println(quiz)
	if check := db.Create(&quiz).Error; check != nil {
		c.Header("access-control-allow-origin", "*")
		c.AbortWithStatus(404)
		fmt.Println(check)
	} else {
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, quiz)
	}
}

func DeleteQuiz(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	id := c.Params.ByName("id")
	var quiz Quiz
	check := db.Where("id = ?", id).Delete(&quiz).Error
	if check != nil {
		c.Header("access-control-allow-origin", "*")
		c.AbortWithStatus(404) //To be decided
		fmt.Println(err)
	}
	c.Header("access-control-allow-origin", "*")
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
