package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"online-shop-backend/internal/app/user/repository"
	"online-shop-backend/internal/domain/dto"
	"online-shop-backend/internal/domain/entity"
	"online-shop-backend/internal/infra/jwt"
)

type UserUsecaseItf interface {
	Register(register dto.Register) error
	Login(login dto.Login) (string, error)
}

type UserUsecase struct {
	userRepository repository.UserMySQLItf
	jwt            jwt.JWTItf
}

func NewUserUsecase(userRepository repository.UserMySQLItf, jwt jwt.JWTItf) UserUsecaseItf {
	return &UserUsecase{
		userRepository: userRepository,
		jwt:            jwt,
	}
}

func (u *UserUsecase) Register(register dto.Register) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		ID:       uuid.New(),
		Name:     register.Name,
		Email:    register.Email,
		Password: string(hashedPassword),
	}

	return u.userRepository.CreateUser(user)
}

func (u *UserUsecase) Login(login dto.Login) (string, error) {
	user, err := u.userRepository.GetUserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "invalid email or password")
	}

	token, err := u.jwt.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return "", err
	}

	return token, nil
}
