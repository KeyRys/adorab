package routes

import (
	"adoend/internal/delivery/http"
	"adoend/internal/repository"
	"adoend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(router *gin.Engine, db *pgx.Conn, secret string) {

	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := http.NewProductHandler(productUsecase)

	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})
}
