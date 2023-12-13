package services

import (
	"fmt"
	"golangapi/controllers/inputs"
	gormdb "golangapi/databases/gorm"
	"golangapi/datalayers"
	"golangapi/models"
)

type IUserService interface {
	CreateUser(registerData inputs.Registration) error
	VerifyUser(loginData inputs.Login) (string, error)
}

type UserService struct {
	UserDataLayer datalayers.UserDatalayer
	CryptoService ICryptoService
	TokenHandler ITokenHandler
}

func NewUserService(userDL datalayers.UserDatalayer, cryptoService ICryptoService, tokenHandler ITokenHandler) IUserService {
	return UserService{
		UserDataLayer: userDL,
		CryptoService: cryptoService,
		TokenHandler: tokenHandler,
	}
}

func NewDefaultUserService() IUserService {
	return UserService{
		UserDataLayer: datalayers.NewGormUserDatalayer(),
		CryptoService: NewDefaultCryptoService(),
		TokenHandler: NewDefaultSymmetricalPasetoTokenHandler(),
	}
}

func (us UserService) CreateUser(registerData inputs.Registration) error {
	hashedPass, err := us.CryptoService.HashPassword(registerData.Password)

	if err != nil {
		return err
	}

	_, err = us.UserDataLayer.CreateUser(
		models.UserInfo{
			Password: hashedPass,
			Email: registerData.Email,
		},
		gormdb.GetDefaultGormClient(),
	)

	return err
}

func (us UserService) VerifyUser(loginData inputs.Login) (string, error) {
	user, err := us.UserDataLayer.FindUserByEmail(loginData.Email, gormdb.GetDefaultGormClient())

	if err != nil {
		fmt.Println("Invalid email")
		return "", fmt.Errorf("invalid combination")
	}

	isOk := us.CryptoService.CheckPasswordHash(loginData.Password, user.Password)

	if !isOk {
		fmt.Println("Invalid pass")
		return "", fmt.Errorf("invalid combination")
	}

	token, err := us.TokenHandler.CreateToken(nil, nil)

	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

