package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"csv-storage/db"

	"csv-storage/config"
	"csv-storage/utils"
	"csv-storage/parsing"

	"csv-storage/service"
)

func main() {

	initiatRedis()

	filePath := os.Args[1]
	go parsing.ParseAndStore(filePath)

	//  create a new *router instance
	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":1321", router))
}

func setUpRedis(rc *config.RedisConfig) {
	url := rc.RedisAdress + ":" + rc.RedisPort
	redisDatabase, err := strconv.Atoi(rc.RedisDataBase)
	redis, err := db.NewRedis(url, rc.RedisPassword, redisDatabase)
	utils.HandleError(err)
	db.SetRepository(*redis)
}

func initiatRedis() {
	r := config.IniatilizeRedisConfig()
	setUpRedis(r)
}
