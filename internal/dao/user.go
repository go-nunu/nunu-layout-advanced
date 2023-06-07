package dao

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserDao struct {
	*Dao
}

func NewUserDao(dao *Dao) *UserDao {
	return &UserDao{
		dao,
	}
}
func (d *UserDao) CreateUser(user *model.User) error {
	if err := d.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (d *UserDao) GetUserById(userId string) (*model.User, error) {
	var user model.User
	if err := d.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return &user, nil
}

func (d *UserDao) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := d.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return &user, nil
}

func (d *UserDao) UpdateUser(user *model.User) error {
	if err := d.db.Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
