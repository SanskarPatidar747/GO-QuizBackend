package genreFunction

import (
	"fmt"
	//"github.com/gin-contrib/cors"                        // Why do we need this package?
	"github.com/gin-gonic/gin" // Using gin as microframework
	"github.com/jinzhu/gorm"   //Using gorm as orm
	_ "modernc.org/sqlite"     //Using modernc.org/sqlite as db driver
)

type Genre struct {
	ID   uint   `json:"id"`
	Name string `gorm:"type:varchar(100);unique_index" json:"name"`
}

/*
	func DataBaseOpener() *gorm.DB{
	    db, err := gorm.Open("sqlite", "./gorm.db")
	    if err != nil {
	        panic("failed to connect table")
	    }
	    defer db.Close()
	    return db
	}
*/
func GetAllGenre(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	var genre []Genre
	if err := db.Find(&genre).Error; err != nil {
		c.Header("access-control-allow-origin", "*")
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, genre)
	}
}

func DeleteGenre(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()
	id := c.Params.ByName("id")
	var genre Genre
	d := db.Where("id = ?", id).Delete(&genre)
	fmt.Println(d)
	c.Header("access-control-allow-origin", "*")
	c.JSON(204, gin.H{"id #" + id: "deleted"})
}

func AddGenre(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()

	var genre Genre
	c.BindJSON(&genre)
	fmt.Println(genre)
	db.Create(&genre)
	c.Header("access-control-allow-origin", "*")
	c.JSON(200, genre)
}

func GetGenre(c *gin.Context) {
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		panic("failed to connect table")
	}
	defer db.Close()

	id := c.Params.ByName("id")
	var genre Genre
	if check := db.Where("id = ?", id).First(&genre).Error; check != nil {
		c.Header("access-control-allow-origin", "*")
		c.AbortWithStatus(404)
		fmt.Println(check)
	} else {
		c.Header("access-control-allow-origin", "*")
		c.JSON(200, genre)
	}
}
