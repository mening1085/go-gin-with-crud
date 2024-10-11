package routes

import (
	"go-crud/controllers"
	"go-crud/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Auth
	r.POST("api/login", controllers.Login)
	r.POST("api/register", controllers.Register)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Route ที่ต้องการการตรวจสอบ JWT
	authApi := r.Group("/api")
	authApi.Use(middleware.AuthMiddleware())
	authApi.GET("/", func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + username})
	})
	// products
	authApi.GET("/products", controllers.GetProducts)
	authApi.GET("/products/:id", controllers.GetProduct)
	authApi.POST("/products", controllers.CreateProduct)
	authApi.PUT("/products/:id", controllers.UpdateProduct)
	authApi.DELETE("/products/:id", controllers.DeleteProduct)

	// categories
	authApi.GET("/categories", controllers.GetCategories)
	authApi.GET("/categories/:id", controllers.GetCategory)
	authApi.POST("/categories", controllers.CreateCategory)
	authApi.PUT("/categories/:id", controllers.UpdateCategory)
	authApi.DELETE("/categories/:id", controllers.DeleteCategory)

	return r
}
