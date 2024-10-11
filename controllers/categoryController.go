package controllers

import (
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryWithProductResponse struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	Products []ProductResponse `json:"products"`
}
type ProductResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	CategoryID uint    `json:"category_id"`
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	models.DB.Preload("Products").Find(&categories)

	var categoryResponse []CategoryWithProductResponse

	for _, category := range categories {
		var productResponse []ProductResponse
		for _, product := range category.Products {
			productResponse = append(productResponse, ProductResponse{
				ID:         product.ID,
				Name:       product.Name,
				Price:      product.Price,
				Quantity:   product.Quantity,
				CategoryID: product.CategoryID,
			})
		}
		categoryResponse = append(categoryResponse, CategoryWithProductResponse{
			ID:       category.ID,
			Name:     category.Name,
			Products: productResponse,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Category",
		"status":  http.StatusOK,
		"data":    categoryResponse,
	})
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := models.DB.Preload("Products").First(&category, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!"})
		return
	}

	var productResponse []ProductResponse
	for _, product := range category.Products {
		productResponse = append(productResponse, ProductResponse{
			ID:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			Quantity:   product.Quantity,
			CategoryID: product.CategoryID,
		})
	}

	categoryResponse := CategoryWithProductResponse{
		ID:       category.ID,
		Name:     category.Name,
		Products: productResponse,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Category",
		"status":  http.StatusOK,
		"data":    categoryResponse,
	})
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&category)
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Category Success",
		"status":  http.StatusOK,
	})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := models.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!"})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&category).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Category Success",
		"status":  http.StatusOK,
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := models.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!"})
		return
	}

	models.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Category Success",
		"status":  http.StatusOK,
	})
}
