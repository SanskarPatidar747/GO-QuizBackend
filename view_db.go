package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "modernc.org/sqlite"
)

// User struct
type User struct {
	ID       uint   `json:"id"`
	UserName string `gorm:"unique_index" json:"user_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

// Genre struct
type Genre struct {
	ID   uint   `json:"id"`
	Name string `gorm:"type:varchar(100);unique_index" json:"name"`
}

// Quiz struct
type Quiz struct {
	ID       uint   `json:"id"`
	Title    string `gorm:"type:varchar(100)" json:"title"`
	Genre_id uint   `json:"genre_id"`
}

// Question struct
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

// Leaderboard struct
type Leaderboard struct {
	ID      uint `json:"id"`
	User_id uint `gorm:"unique_index:idx_name_code" json:"user_id"`
	Quiz_id uint `gorm:"unique_index:idx_name_code" json:"quiz_id"`
	Score   int  `json:"score"`
}

func main() {
	// Open database
	db, err := gorm.Open("sqlite", "./gorm.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("🗄️  Database Connected Successfully!")
	fmt.Println("📊 Viewing Database Contents...\n")

	// Check if tables exist and show data
	var users []User
	var genres []Genre
	var quizzes []Quiz
	var questions []Question
	var leaderboards []Leaderboard

	// Get Users
	if db.HasTable(&User{}) {
		db.Find(&users)
		fmt.Printf("👥 USERS (%d records):\n", len(users))
		for _, user := range users {
			fmt.Printf("  ID: %d | Username: %s | Name: %s | Email: %s | Admin: %t\n", 
				user.ID, user.UserName, user.Name, user.Email, user.Admin)
		}
		fmt.Println()
	} else {
		fmt.Println("❌ Users table not found")
	}

	// Get Genres
	if db.HasTable(&Genre{}) {
		db.Find(&genres)
		fmt.Printf("📚 GENRES (%d records):\n", len(genres))
		for _, genre := range genres {
			fmt.Printf("  ID: %d | Name: %s\n", genre.ID, genre.Name)
		}
		fmt.Println()
	} else {
		fmt.Println("❌ Genres table not found")
	}

	// Get Quizzes
	if db.HasTable(&Quiz{}) {
		db.Find(&quizzes)
		fmt.Printf("🎯 QUIZZES (%d records):\n", len(quizzes))
		for _, quiz := range quizzes {
			fmt.Printf("  ID: %d | Title: %s | Genre ID: %d\n", 
				quiz.ID, quiz.Title, quiz.Genre_id)
		}
		fmt.Println()
	} else {
		fmt.Println("❌ Quizzes table not found")
	}

	// Get Questions
	if db.HasTable(&Question{}) {
		db.Find(&questions)
		fmt.Printf("❓ QUESTIONS (%d records):\n", len(questions))
		for _, question := range questions {
			fmt.Printf("  ID: %d | Question: %s | Quiz ID: %d | Score: %d\n", 
				question.ID, question.Question[:min(50, len(question.Question))], question.Quiz_id, question.Score)
		}
		fmt.Println()
	} else {
		fmt.Println("❌ Questions table not found")
	}

	// Get Leaderboards
	if db.HasTable(&Leaderboard{}) {
		db.Find(&leaderboards)
		fmt.Printf("🏆 LEADERBOARDS (%d records):\n", len(leaderboards))
		for _, lb := range leaderboards {
			fmt.Printf("  ID: %d | User ID: %d | Quiz ID: %d | Score: %d\n", 
				lb.ID, lb.User_id, lb.Quiz_id, lb.Score)
		}
		fmt.Println()
	} else {
		fmt.Println("❌ Leaderboards table not found")
	}

	// Show table info
	fmt.Println("📋 DATABASE INFO:")
	fmt.Printf("  Database file: gorm.db\n")
	fmt.Printf("  Total tables: 5\n")
	fmt.Printf("  Total records: %d\n", len(users)+len(genres)+len(quizzes)+len(questions)+len(leaderboards))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
} 