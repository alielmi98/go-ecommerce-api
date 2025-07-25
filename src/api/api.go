package api

import (
	"fmt"
	"log"

	"github.com/alielmi98/go-ecommerce-api/api/middlewares"
	"github.com/alielmi98/go-ecommerce-api/api/routers"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/docs"
	"github.com/alielmi98/go-ecommerce-api/events"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(middlewares.Cors(cfg))
	dispatcher := events.NewEventDispatcher()
	RegisterRoutes(r, cfg, dispatcher)
	RegisterSwagger(r, cfg)
	log.Printf("Caller:%s Level:%s Msg:%s", constants.General, constants.Startup, "Started")
	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))

}

func RegisterRoutes(r *gin.Engine, cfg *config.Config, dispatcher *events.EventDispatcher) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		//Account
		account := v1.Group("/account")
		routers.Account(account, cfg)
		//Token
		token := v1.Group("/token")
		routers.Token(token, cfg)

		//Admin Panel Routers
		shop := v1.Group("/shop")
		{

			//category
			categories := shop.Group("/categories")
			routers.Category(categories, cfg)
			//file
			files := shop.Group("/files")
			routers.File(files, cfg)
			//product
			products := shop.Group("/products")
			routers.Product(products, cfg, dispatcher)
			//productImage
			productImages := shop.Group("/product-images")
			routers.ProductImage(productImages, cfg)
			//productReview
			productReview := shop.Group("/product-reviews")
			routers.ProductReview(productReview, cfg)

			//cartItem
			cartItem := shop.Group("/cart-items")
			routers.CartItem(cartItem, cfg, dispatcher)
			//cart
			cart := shop.Group("/carts")
			routers.Cart(cart, cfg)

			//orderItem
			orderItem := shop.Group("/order-items")
			routers.OrderItem(orderItem, cfg)
			//order
			order := shop.Group("/orders")
			routers.Order(order, cfg)

			//checkout
			checkout := shop.Group("/checkout")
			routers.CheckOutHandler(checkout, cfg)

			//payment
			payment := shop.Group("/payments")
			routers.Payment(payment, cfg)

		}

	}

}
func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang URL Ecommerce Service api documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.InternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
