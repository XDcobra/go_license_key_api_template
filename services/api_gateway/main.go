package main

import (
	"context"
	"fmt"
	"github.com/XDcobra/go_license_key_api_template/database"
	"log"
	"time"
)

// Main function
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())

	// create connection to redis database
	redisClient := database.ConnectionRedisDB()

	// check connection
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to Redis")
	}
}
