package bootstrap

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	producthandler "online-shop-backend/internal/app/product/interface/rest"
	"online-shop-backend/internal/app/product/repository"
	productusecase "online-shop-backend/internal/app/product/usecase"
	"online-shop-backend/internal/infra/env"
	"online-shop-backend/internal/infra/fiber"
	"online-shop-backend/internal/infra/mysql"
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

	app := fiber.New()
	v1 := app.Group("/api/v1")

	productRepository := repository.NewProductMySQL(database)
	productUsecase := productusecase.NewProductUsecase(productRepository)
	producthandler.NewProductHandler(v1, productUsecase, val)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
