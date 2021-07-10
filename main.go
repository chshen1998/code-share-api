package main

import (
	"database/sql"
	"fmt"
	"gin_project/config"
	"gin_project/routers"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-18-235-4-83.compute-1.amazonaws.com"
	port     = 5432
	user     = "dacjdwnnxlvkev"
	password = "10bd3a55e52242118a1c45b314f93c1ef0e62be94c2fe6374c0077e77e97d24d"
	dbname   = "d7ulng2l63hcrc"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	config.DB, _ = sql.Open("postgres", connStr)
	defer config.DB.Close()

	r := routers.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
