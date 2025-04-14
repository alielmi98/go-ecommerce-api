package usecase

import (
	"log"

	"github.com/alielmi98/go-ecommerce-api/api/dto"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/domin/models"
	"github.com/alielmi98/go-ecommerce-api/domin/repository"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	cfg          *config.Config
	repo         repository.UserRepository
	tokenUsecase *TokenUsecase
}

func NewUserUsecase(cfg *config.Config, repository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		cfg:          cfg,
		repo:         repository,
		tokenUsecase: NewTokenUsecase(cfg),
	}
}

// Register by username
func (s *UserUsecase) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{Username: req.Username, FirstName: req.FirstName, LastName: req.LastName, Email: req.Email}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.General, constants.HashPassword, err.Error())
		return err
	}
	u.Password = string(hp)

	err = s.repo.Create(&u)
	if err != nil {
		return err
	}
	return nil

}

// LoginByUsername
func (s *UserUsecase) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UsernameOrPasswordInvalid}
	}

	tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Username: user.Username, Email: user.Email, MobileNumber: user.MobileNumber}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}

	token, err := s.tokenUsecase.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil
}
