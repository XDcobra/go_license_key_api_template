package main

import (
	"fmt"
	dummycontroller "github.com/XDcobra/go_license_key_api_template/controller/DummyController"
	mysqlcontroller "github.com/XDcobra/go_license_key_api_template/controller/MySQLController"
	rediscontroller "github.com/XDcobra/go_license_key_api_template/controller/RedisController"
	dbmanagementMySQL "github.com/XDcobra/go_license_key_api_template/database/MySQL"
	dbmanagementRedis "github.com/XDcobra/go_license_key_api_template/database/Redis"
	router "github.com/XDcobra/go_license_key_api_template/router"
	"log"
)

// Main function
func main() {
	// create connection to redis database
	redisClient := dbmanagementRedis.ConnectionRedisDB()

	// create connection to mysql database using GORM
	mysqlClient := dbmanagementMySQL.ConnectionMySQLDB()

	// do mysql model automigration
	err := dbmanagementMySQL.Automigration(mysqlClient)
	if err != nil {
		log.Fatalf("Error while migrating MySQL Database: %v", err)
	}

	// create server / register middlewares
	app := router.CreateServer()

	// create controllers
	dummyController := dummycontroller.NewDummyController()
	redisController := rediscontroller.NewRedisController(redisClient)
	mysqlController := mysqlcontroller.NewMySQLController(mysqlClient)

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
