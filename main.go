package main

import (
	"gin_project/config"
	"gin_project/routers"

	_ "github.com/lib/pq"
)

func main() {
	config.InitDB()
	config.InitRedis()
	defer config.DB.Close()

	r := routers.InitRouter()
	r.Run() // listen and serve on localhost:8080
}
