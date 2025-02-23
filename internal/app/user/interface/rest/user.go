package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"online-shop-backend/internal/app/user/usecase"
	"online-shop-backend/internal/domain/dto"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecaseItf
	Validator   *validator.Validate
}

func NewUserHandler(
	routerGroup fiber.Router,
	userUsecase usecase.UserUsecaseItf,
	validator *validator.Validate,
) {
	UserHandler := UserHandler{
		UserUsecase: userUsecase,
		Validator:   validator,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", UserHandler.RegisterUser)
	routerGroup.Post("/login", UserHandler.LoginUser)
}

func (u *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	var register dto.Register

	if err := ctx.BodyParser(&register); err != nil {
		return err
	}

	if err := u.Validator.Struct(register); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := u.UserUsecase.Register(register); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Register Success",
	})
}

func (u *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	var login dto.Login

	if err := ctx.BodyParser(&login); err != nil {
		return err
	}

	if err := u.Validator.Struct(login); err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	}

	token, err := u.UserUsecase.Login(login)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Login Success",
		"payload": fiber.Map{
			"token": token,
		},
	})
}
