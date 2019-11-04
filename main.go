package main

import (
	"blog/models"
	"blog/server"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.InitDB()

	log.Fatal(server.Started())

	models.CloseDB()
}
