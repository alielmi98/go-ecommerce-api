package usecase

import (
	"log"

	"github.com/alielmi98/go-ecommerce-api/api/dto"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/data/db/repository"
	"github.com/alielmi98/go-ecommerce-api/data/models"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	cfg  *config.Config
	repo *repository.UserRepository
}

func NewUserUsecase(cfg *config.Config) *UserUsecase {
	return &UserUsecase{
		cfg:  cfg,
		repo: repository.NewUserRepository(),
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
