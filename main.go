package main

import (
	"gin_project/config"
	"gin_project/routers"

	_ "github.com/lib/pq"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	r := routers.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
