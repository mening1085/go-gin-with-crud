package controllers

import (
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProductWithCategoryResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	CategoryID uint    `json:"category_id"`
	Category   CategoryResponse
}

func GetProducts(c *gin.Context) {

	var products []models.Product
	models.DB.Preload("Category").Find(&products)

	var productResponse []ProductWithCategoryResponse

	for _, product := range products {
		productResponse = append(productResponse, ProductWithCategoryResponse{
			ID:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			Quantity:   product.Quantity,
			CategoryID: product.CategoryID,
			Category: CategoryResponse{
				ID:   product.Category.ID,
				Name: product.Category.Name,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetProducts",
		"status":  http.StatusOK,
		"data":    productResponse,
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := models.DB.Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}

	productResponse := ProductWithCategoryResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Quantity:   product.Quantity,
		CategoryID: product.CategoryID,
		Category: CategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GetProduct",
		"status":  http.StatusOK,
		"data":    productResponse,
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Product Success",
		"status":  http.StatusOK,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product Success",
		"status":  http.StatusOK,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Product Success",
		"status":  http.StatusOK,
	})
}
