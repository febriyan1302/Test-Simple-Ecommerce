// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package router

import (
	"github.com/febriyan1302/catalog-service/internal/controller"
	"github.com/febriyan1302/catalog-service/internal/repository"
	"github.com/febriyan1302/catalog-service/internal/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from wire.go:

// ping
func InitPing(db *gorm.DB) controller.Controller {
	repositoryRepository := repository.RegisterRepository(db)
	serviceService := service.RegisterService(repositoryRepository)
	controllerController := controller.RegisterController(serviceService)
	return controllerController
}

func InitAuth(db *gorm.DB) controller.Auth {
	repositoryAuthRepository := repository.ConstructAuthRepository(db)
	serviceAuthService := service.ConstructAuthService(repositoryAuthRepository)
	controllerAuthController := controller.ConstructAuthController(serviceAuthService)
	return controllerAuthController
}

func InitUser(db *gorm.DB) controller.User {
	userRepository := repository.ConstructUserRepository(db)
	userService := service.ConstructUserService(userRepository)
	userController := controller.ConstructUserController(userService)
	return userController
}

func InitProduct(db *gorm.DB) controller.Product {
	repositoryProductRepository := repository.ConstructProductRepopsitory(db)
	serviceProductService := service.ConstructProductService(repositoryProductRepository)
	controllerProductController := controller.ConstructorProductController(serviceProductService)
	return controllerProductController
}

func InitCategory(db *gorm.DB) controller.Category {
	repositoryCategoryRepository := repository.ConstructorCategoryRepository(db)
	serviceCategoryService := service.ConstructorCategoryService(repositoryCategoryRepository)
	controllerCategoryController := controller.ConstructorCategoryController(serviceCategoryService)
	return controllerCategoryController
}

func InitCart(db *gorm.DB) controller.Cart {
	repositoryCartRepository := repository.ConstructorCartRepository(db)
	serviceCartService := service.ConstructorCartService(repositoryCartRepository)
	controllerCartController := controller.ConstructorCartController(serviceCartService)
	return controllerCartController
}

// wire.go:

// auth
var (
	authRepository = wire.NewSet(repository.ConstructAuthRepository, wire.Bind(
		new(repository.Auth),
		new(*repository.AuthRepository),
	),
	)

	authService = wire.NewSet(service.ConstructAuthService, wire.Bind(
		new(service.Auth),
		new(*service.AuthService),
	),
	)

	authController = wire.NewSet(controller.ConstructAuthController, wire.Bind(
		new(controller.Auth),
		new(*controller.AuthController),
	),
	)
)

// users
var (
	usersRepository = wire.NewSet(repository.ConstructUserRepository, wire.Bind(
		new(repository.User),
		new(*repository.UserRepository),
	),
	)

	usersService = wire.NewSet(service.ConstructUserService, wire.Bind(
		new(service.User),
		new(*service.UserService),
	),
	)

	usersController = wire.NewSet(controller.ConstructUserController, wire.Bind(
		new(controller.User),
		new(*controller.UserController),
	),
	)
)

// product
var (
	productRepository = wire.NewSet(repository.ConstructProductRepopsitory, wire.Bind(
		new(repository.Product),
		new(*repository.ProductRepository),
	),
	)

	productService = wire.NewSet(service.ConstructProductService, wire.Bind(
		new(service.Product),
		new(*service.ProductService),
	),
	)

	productController = wire.NewSet(controller.ConstructorProductController, wire.Bind(
		new(controller.Product),
		new(*controller.ProductController),
	),
	)
)

// category
var (
	categoryRepository = wire.NewSet(repository.ConstructorCategoryRepository, wire.Bind(
		new(repository.Category),
		new(*repository.CategoryRepository),
	),
	)

	categoryService = wire.NewSet(service.ConstructorCategoryService, wire.Bind(
		new(service.Category),
		new(*service.CategoryService),
	),
	)

	categoryController = wire.NewSet(controller.ConstructorCategoryController, wire.Bind(
		new(controller.Category),
		new(*controller.CategoryController),
	),
	)
)

// cart
var (
	cartRepository = wire.NewSet(repository.ConstructorCartRepository, wire.Bind(
		new(repository.Cart),
		new(*repository.CartRepository),
	),
	)

	cartService = wire.NewSet(service.ConstructorCartService, wire.Bind(
		new(service.Cart),
		new(*service.CartService),
	),
	)

	cartController = wire.NewSet(controller.ConstructorCartController, wire.Bind(
		new(controller.Cart),
		new(*controller.CartController),
	),
	)
)
