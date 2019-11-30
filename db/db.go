package db

import (
	"fmt"
	"github.com/X1Zeth2X/projectify-api/models"
	"github.com/jinzhu/gorm"
	"os"
	// Import postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load() // Load .env
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	// Change sslmode from disable -> enable when running in prod.
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	/*
		Keep Automigrate commented, only use it to make migrations.
		You wouldn't want to keep performing migrations when running the app.
	*/
	db.Debug().AutoMigrate(
		&models.User{},
		&models.Project{},
	)

}

// GetDB Returns the current db session
func GetDB() *gorm.DB {
	return db
}
