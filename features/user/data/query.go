package data

import (
	"capstone-project/features/user"
	"capstone-project/middlewares"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type dataUser struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &dataUser{
		db,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (repo *dataUser) LoginUser(data user.UserCore) (token string, err error) {
	var user User
	tx := repo.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&user)
	result := CheckPasswordHash(data.Password, user.Password)
	if result == false {
		return "", errors.New("password is incorrect")
	}

	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", errors.New("acoount not found")
	}

	token, err = middlewares.CreateToken(int(user.ID), user.Role, user.Nama_User)
	if err != nil {
		return "", err
	}
	return token, err
}
