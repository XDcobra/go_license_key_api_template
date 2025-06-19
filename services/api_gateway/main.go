package main

import (
	"context"
	"fmt"
	dummycontroller "github.com/XDcobra/go_license_key_api_template/controller/DummyController"
	rediscontroller "github.com/XDcobra/go_license_key_api_template/controller/RedisController"
	"github.com/XDcobra/go_license_key_api_template/database"
	router "github.com/XDcobra/go_license_key_api_template/router"
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

	// create server / register middlewares
	app := router.CreateServer()

	// create controllers
	dummyController := dummycontroller.NewDummyController()
	redisController := rediscontroller.NewRedisController(redisClient)

	// register routes
	routerClient := router.RegisterRoutes(app, redisController, dummyController)

	// start http server
	err := routerClient.Listen(":8000")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listening on 8000")
	}

}
