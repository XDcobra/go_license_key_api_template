package main

import (
	"context"
	"fmt"
	dummycontroller "github.com/XDcobra/go_license_key_api_template/controller/DummyController"
	mysqlcontroller "github.com/XDcobra/go_license_key_api_template/controller/MySQLController"
	rediscontroller "github.com/XDcobra/go_license_key_api_template/controller/RedisController"
	"github.com/XDcobra/go_license_key_api_template/database"
	router "github.com/XDcobra/go_license_key_api_template/router"
	mysqlservice "github.com/XDcobra/go_license_key_api_template/services/MySQL"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// Main function
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create connection to redis database
	redisClient := database.ConnectionRedisDB()

	// check connection
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to Redis")
	}

	// create connection to mysql database using GORM
	dsn := "user:password@tcp(mysql:3306)/example_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to MySQL Database: %v", err)
	}

	// do automigration
	err = mysqlservice.Automigration(db)
	if err != nil {
		log.Fatalf("Error while migrating MySQL Database: %v", err)
	}

	// create server / register middlewares
	app := router.CreateServer()

	// create controllers
	dummyController := dummycontroller.NewDummyController()
	redisController := rediscontroller.NewRedisController(redisClient)
	mysqlController := mysqlcontroller.NewMySQLController(db)

	// register routes
	routerClient := router.RegisterRoutes(app, redisController, dummyController, mysqlController)

	// start http server
	err = routerClient.Listen(":8000")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listening on 8000")
	}

}
