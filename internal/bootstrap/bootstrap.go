package bootstrap

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	producthandler "online-shop-backend/internal/app/product/interface/rest"
	productrepository "online-shop-backend/internal/app/product/repository"
	productusecase "online-shop-backend/internal/app/product/usecase"
	userhandler "online-shop-backend/internal/app/user/interface/rest"
	userrepository "online-shop-backend/internal/app/user/repository"
	userusecase "online-shop-backend/internal/app/user/usecase"
	"online-shop-backend/internal/infra/env"
	"online-shop-backend/internal/infra/fiber"
	"online-shop-backend/internal/infra/jwt"
	"online-shop-backend/internal/infra/mysql"
	"online-shop-backend/internal/middleware"
	"time"
)

func Start() error {
	config, err := env.New()
	if err != nil {
		return err
	}

	database, err := mysql.New(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName),
	)
	if err != nil {
		return err
	}

	if err := mysql.Migrate(database); err != nil {
		return err
	}

	val := validator.New()

	jwtExpiredTime, err := time.ParseDuration(config.JWTTTL)
	if err != nil {
		return err
	}
	jwtInstance := jwt.NewJWT(config.JWTSecret, jwtExpiredTime)

	app := fiber.New()
	v1 := app.Group("/api/v1")

	userRepository := userrepository.NewUserMySQL(database)
	productRepository := productrepository.NewProductMySQL(database)

	userUsecase := userusecase.NewUserUsecase(userRepository, jwtInstance)
	productUsecase := productusecase.NewProductUsecase(productRepository)

	midw := middleware.NewMiddleware(jwtInstance)

	userhandler.NewUserHandler(v1, userUsecase, val)
	producthandler.NewProductHandler(v1, productUsecase, midw, val)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
