package main

import (
	"gin_project/config"
	"gin_project/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	r := routers.InitRouter()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Run()
}
