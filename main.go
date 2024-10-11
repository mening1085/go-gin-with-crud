package main

import (
	"go-crud/models"
	"go-crud/routes"
)

func main() {
	models.InitDB() // เรียกใช้งานการเชื่อมต่อฐานข้อมูล

	r := routes.SetupRouter()
	r.Run()
}
