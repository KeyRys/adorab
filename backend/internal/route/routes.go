package routes

import (
	"adoend/internal/delivery/http"
	"adoend/internal/delivery/http/middleware"
	"adoend/internal/repository"
	"adoend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(router *gin.Engine, db *pgx.Conn, secret string) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is running",
		})
	})

	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := http.NewProductHandler(productUsecase)

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo, secret)
	authHandler := http.NewAuthHandler(authUsecase)

	cartRepo := repository.NewCartRepository(db)
	cartUsecase := usecase.NewCartUsecase(cartRepo)
	cartHandler := http.NewCartHandler(cartUsecase)

	profileRepo := repository.NewProfileRepository(db)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)
	profileHandler := http.NewProfileHandler(profileUsecase)

	sellerRepo := repository.NewSellerRepository(db)
	sellerUsecase := usecase.NewSellerUsecase(sellerRepo)
	sellerHandler := http.NewSellerHandler(sellerUsecase)

	checkoutRepo := repository.NewCheckoutRepository(db)
	checkoutUsecase := usecase.NewCheckoutUsecase(checkoutRepo)
	checkoutHandler := http.NewCheckoutHandler(checkoutUsecase)

	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := http.NewOrderHandler(orderUsecase)
	sellerOrderUsecase := usecase.NewSellerOrderUsecase(orderRepo)
	sellerOrderHandler := http.NewSellerOrderHandler(sellerOrderUsecase)

	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProductByID)

	auth := router.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(secret))

	protected.POST("/cart/add", cartHandler.AddToCart)
	protected.GET("/cart", cartHandler.GetCart)
	protected.DELETE("/cart/item/:id", cartHandler.RemoveItem)

	protected.GET("/profile/me", profileHandler.GetMyProfile)
	protected.PUT("/profile", profileHandler.UpdateProfile)

	protected.POST("/seller/up", sellerHandler.BecomeSeller)
	protected.GET("/seller/rabbits", sellerHandler.GetSellerRabbits)
	protected.POST("/seller/rabbits", sellerHandler.CreateRabbit)
	protected.PUT("/seller/rabbits/:id", sellerHandler.UpdateRabbit)
	protected.DELETE("/seller/rabbits/:id", sellerHandler.DeleteRabbit)

	protected.POST("/checkout", checkoutHandler.Checkout)

	protected.GET("/orders/me", orderHandler.GetBuyerOrders)
	protected.GET("/orders/seller", sellerOrderHandler.GetSellerOrders)
	protected.PUT("/orders/:id", sellerOrderHandler.UpdateOrderStatus)
}
