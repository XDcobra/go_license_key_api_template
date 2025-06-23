package MySQL

import (
	"fmt"
	MySQlModels "github.com/XDcobra/gofiber-starter-stack/model/MySQL"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectionMySQLDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to MySQL Database: %v", err)
	}

	// check mysql database connection
	sqlDB, err := db.DB() // get underlying *sql.DB
	if err != nil {
		log.Fatalf("Unable to get raw DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	} else {
		log.Println("Connected to MySQL")
	}

	return db
}

func Automigration(db *gorm.DB) error {
	// Auto-migrate User (and more models if added in the future)
	err := db.AutoMigrate(&MySQlModels.User{})

	return err
}
