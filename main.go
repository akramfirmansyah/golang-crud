package main

import (
	"github.com/akramfirmansyah/go-crud/controllers"
	"github.com/akramfirmansyah/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDatabase()
	initializers.Migrate()
}

func main() {
	r := gin.Default()

	mahasiswaAPI := r.Group("/api/mahasiswa")
	{
		mahasiswaAPI.POST("/", controllers.MahasiswaCreate)
		mahasiswaAPI.GET("/", controllers.MahasiswaGet)
		mahasiswaAPI.GET("/:nim", controllers.MahasiswaGetByNim)
		mahasiswaAPI.PUT("/:nim", controllers.MahasiswaUpdate)
		mahasiswaAPI.DELETE("/:nim", controllers.MahasiswaDelete)
	}

	r.Run()
}
