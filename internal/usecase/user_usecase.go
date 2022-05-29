package usecase

import (
	"aulia-majoo-test/internal/domain"
	"aulia-majoo-test/internal/dto"
	"aulia-majoo-test/internal/models"
	"aulia-majoo-test/pkg/util"
	"errors"
)

type UserUsecase struct {
	uRepo domain.UserRepository
}

func NewUserUsecase(uRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{uRepo: uRepo}
}

func (r *UserUsecase) UserLogin(req dto.LoginRequestDTO) (models.Responses, error) {

	// get merchant by username
	res, err := r.uRepo.GetUserByUsername(req.Username)
	if err != nil {
		return models.Responses{}, errors.New("user not found")
	}
	// compare password
	checkPassword := util.CheckPassword(req.Password, res.Password)
	if checkPassword != nil {
		return models.Responses{}, errors.New("password invalid")

	}

	// Generate Auth Token
	token, err := util.CreateToken(res.ID)
	if err != nil {
		return models.Responses{}, errors.New("Cannot generate token")
	}
	loginData := models.LoginResponse{
		Token: models.Token{
			TokenType:   "Bearer",
			AccessToken: token,
		},
		Username:     res.Username,
		Name:         res.Name,
		MerchantName: res.Merchants.MerchantName,
	}

	// return Response
	return models.SuccessResponses(loginData), nil
}

func (r *UserUsecase) UserRegister(req dto.RegisterRequestDTO) (models.Responses, error) {
	password, err := util.HashPassword(req.Password)
	if err != nil {
		return models.InternalServerResponses("Internal Server Error"), err
	}

	checkUsername, err := r.uRepo.GetUserByUsername(req.Username)
	if checkUsername != nil {
		return models.BadRequestResponses("Username is Already Registered"), nil
	}
	if err != nil {
		return models.InternalServerResponses("Internal Server Error"), err
	}

	requestUser := &domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: password,
		Merchants: domain.Merchant{
			MerchantName: req.MerchantName,
		},
	}
	user, err := r.uRepo.CreateUser(requestUser)
	if err != nil {
		return models.InternalServerResponses("Internal Server Error"), err
	}

	token, tokenErr := util.CreateToken(user.ID)
	if tokenErr != nil {
		return models.InternalServerResponses("Internal Server Error"), err
	}

	// get merchant by username
	res, err := r.uRepo.GetUserByUsername(req.Username)
	if err != nil {
		return models.Responses{}, errors.New("user not found")
	}

	loginData := models.RegisterResponse{
		Token: models.Token{
			TokenType:   "Bearer",
			AccessToken: token,
		},
		Username:     res.Username,
		Name:         res.Name,
		MerchantName: res.Merchants.MerchantName,
	}

	// return Response
	return models.SuccessResponses(loginData), nil
}
