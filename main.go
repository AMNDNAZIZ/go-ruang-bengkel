package main

import (
	"github.com/AMNDNAZIZ/go-ruang-bengkel/controllers/usercontroller"
	"github.com/AMNDNAZIZ/go-ruang-bengkel/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/user/get-data", usercontroller.Index)
	r.GET("/api/user/get-data/:id", usercontroller.Show)
	r.POST("/api/user/create", usercontroller.Create)
	r.PUT("/api/user/create/:id", usercontroller.Update)
	r.DELETE("/api/user/delete-user", usercontroller.Delete)

	r.Run()
}
