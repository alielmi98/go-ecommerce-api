package repository

import (
	"errors"
	"log"

	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/data/db"
	"github.com/alielmi98/go-ecommerce-api/data/models"
	"github.com/alielmi98/go-ecommerce-api/pkg/service_errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.GetDb()}
}

func (r *UserRepository) Create(user *models.User) error {
	exists, err := r.existsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = r.existsByUsername(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}
	roleId, err := r.getDefaultRole()
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.DefaultRoleName, err.Error())
		return err
	}
	tx := r.db.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Rollback, err.Error())
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: user.Id}).Error
	if err != nil {
		tx.Rollback()
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Rollback, err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.
		Model(&models.User{}).
		Where("username = ?", username).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UsernameOrPasswordInvalid}
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := r.db.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return false, err
	}
	return exists, nil
}

func (r *UserRepository) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := r.db.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Select, err.Error())
		return false, err
	}
	return exists, nil
}

func (r *UserRepository) getDefaultRole() (roleId int, err error) {

	if err = r.db.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
