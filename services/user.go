package services

import (
	"fmt"
	"golangapi/controllers/inputs"
	"golangapi/controllers/outputs"
	gormdb "golangapi/databases/gorm"
	"golangapi/datalayers"
	"golangapi/models"

	"gorm.io/gorm"
)

type IUserService interface {
	CreateUser(registerData inputs.Registration) error
	VerifyUser(loginData inputs.Login) (string, error)

	SoftDeleteUser(uuid string) error
	HardDeleteUser(uuid string) error

	GetProfile(uuid string) (*outputs.UserProfile, error)
}

type UserService struct {
	UserDataLayer datalayers.UserDatalayer
	CryptoService ICryptoService
	TokenHandler  ITokenHandler
	GormDb        *gorm.DB
}

func NewUserService(
	userDL datalayers.UserDatalayer,
	cryptoService ICryptoService,
	tokenHandler ITokenHandler,
	gormDb *gorm.DB,
) IUserService {
	return UserService{
		UserDataLayer: userDL,
		CryptoService: cryptoService,
		TokenHandler:  tokenHandler,
		GormDb:        gormDb,
	}
}

func NewDefaultUserService() IUserService {
	return UserService{
		UserDataLayer: datalayers.NewGormUserDatalayer(),
		CryptoService: NewDefaultCryptoService(),
		TokenHandler:  NewDefaultSymmetricalPasetoTokenHandler(),
		GormDb:        gormdb.GetDefaultGormClient(),
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
			Email:    registerData.Email,
		},
		us.GormDb,
	)

	return err
}

func (us UserService) VerifyUser(loginData inputs.Login) (string, error) {
	user, err := us.UserDataLayer.GetHashedUserPassword(loginData.Email, us.GormDb)

	if err != nil {
		fmt.Println("Invalid email")
		return "", fmt.Errorf("invalid combination")
	}

	isOk := us.CryptoService.CheckPasswordHash(loginData.Password, user.Password)

	if !isOk {
		fmt.Println("Invalid pass")
		return "", fmt.Errorf("invalid combination")
	}

	token, err := us.TokenHandler.CreateToken(map[string]string{
		"uuid": user.Uuid.String(),
	}, nil)

	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

func (us UserService) SoftDeleteUser(uuid string) error {
	return us.UserDataLayer.SoftDeleteUser(uuid, us.GormDb)
}

func (us UserService) HardDeleteUser(uuid string) error {
	return us.UserDataLayer.HardDeleteUser(uuid, us.GormDb)
}

func (us UserService) GetProfile(uuid string) (*outputs.UserProfile, error) {
	userInfo, err := us.UserDataLayer.FindUserByUuid(uuid, us.GormDb)

	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
