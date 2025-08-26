package questionFunction

import (
	"fmt"
	//"github.com/gin-contrib/cors"                        // Why do we need this package?
	"github.com/gin-gonic/gin" // Using gin as microframework
	"github.com/jinzhu/gorm"   //Using gorm as orm
	_ "modernc.org/sqlite"     //Using modernc.org/sqlite as db driver
)

type Question struct {
	ID       uint   `json:"id"`
	Question string `json:"question"`
	Answer   string `gorm:"type:varchar(4)" json:"answer"`
	Quiz_id  uint   `json:"quiz_id"`
	Multi    bool   `json:"multi"`
	Score    int    `json:"score"`
	Option_a string `json:"option_a"`
	Option_b string `json:"option_b"`
	Option_c string `json:"option_c"`
	Option_d string `json:"option_d"`
}

func DataBaseOpener() *gorm.DB {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	return db
}

func UpdateQuestion(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()

	var question Question
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&question).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&question)
	db.Save(&question)
	c.Header("access-control-allow-origin", "*")
	c.JSON(200, question)
}

func GetAllQuestions(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	quiz_id := c.Params.ByName("quiz_id")
	var question []Question
	if check := db.Where("quiz_id = ?", quiz_id).Find(&question).Error; check != nil {
		c.AbortWithStatus(404)
		fmt.Println(check)
	} else {
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, question)
	}
}

func GetQuestion(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	id := c.Params.ByName("id")
	var question Question
	if check := db.Where("id = ?", id).First(&question).Error; check != nil {
		c.Header("access-control-allow-origin", "*")
		c.AbortWithStatus(404)
		fmt.Println(check)
	} else {
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, question)
	}
}

func AddQuestion(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()

	var question Question
	c.BindJSON(&question)
	fmt.Println(question)
	db.Create(&question)
	c.Header("access-control-allow-origin", "*")
	c.JSON(200, question)
}

func DeleteQuestion(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	id := c.Params.ByName("id")
	var question Question
	check := db.Where("id = ?", id).Delete(&question)
	if check != nil {
		c.Header("access-control-allow-origin", "*")
		c.AbortWithStatus(404) //To be decided
		fmt.Println(err)
	}
	c.Header("access-control-allow-origin", "*")
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
