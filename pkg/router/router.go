package router

import (
	"github.com/febriyan1302/catalog-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteApi(app *fiber.App, db *gorm.DB, jwt *fiber.Handler) {

	diPing := InitPing(db)
	app.Get("/ping", diPing.PingController)

	api := app.Group("/api")

	diAuth := InitAuth(db)
	api.Post("/auth", diAuth.AuthController)

	diUser := InitUser(db)
	api.Post("/register", diUser.CreateController)

	diProduct := InitProduct(db)
	product := api.Group("/product")
	product.Get("/", middleware.DeserializeUser, diProduct.ListProduct)
	product.Get("/detail/:id", middleware.DeserializeUser, diProduct.DetailProduct)

	diCategory := InitCategory(db)
	api.Get("/category", middleware.DeserializeUser, diCategory.ListCategory)

	diCart := InitCart(db)
	cart := api.Group("/cart")
	cart.Post("/", middleware.DeserializeUser, diCart.Create)
	cart.Get("/", middleware.DeserializeUser, diCart.Get)
}
