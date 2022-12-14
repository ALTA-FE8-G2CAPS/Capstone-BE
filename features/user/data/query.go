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
	tx := repo.db.Where("email = ?", data.Email).First(&user)
	result := CheckPasswordHash(data.Password, user.Password)
	if !result {
		return "", errors.New("password is incorrect")
	}

	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", errors.New("acount not found")
	}

	token, err = middlewares.CreateToken(int(user.ID), user.Role, user.Name_User, user.Foto_user, user.User_owner)
	if err != nil {
		return "", err
	}
	return token, err
}

func (repo *dataUser) InsertData(data user.UserCore) (row int, err error) {
	var user User
	hash, err := HashPassword(data.Password)
	if err != nil {
		return -1, err
	}
	user.Name_User = data.Name_User
	user.Email = data.Email
	user.Password = hash
	user.Role = data.Role
	user.Address_user = data.Address_user
	user.Foto_user = data.Foto_user
	user.User_owner = data.User_owner

	tx := repo.db.Create(&user)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataUser) SelectAllUser() ([]user.UserCore, error) {
	var users []User
	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	userCore := toCoreList(users)
	return userCore, nil
}

func (repo *dataUser) SelectUserById(id int) (user.UserCore, error) {
	var userList User
	userList.ID = uint(id)
	tx := repo.db.Where("id = ?", id).First(&userList)
	if tx.Error != nil {
		return user.UserCore{}, tx.Error
	}
	userData := userList.toCore()
	return userData, nil
}

func (repo *dataUser) DeleteUser(id int) (row int, err error) {
	var user User
	user.ID = uint(id)
	tx := repo.db.Unscoped().Delete(&user)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataUser) UpdateUser(data user.UserCore) (int, error) {
	var userUpdate User
	txDataOld := repo.db.First(&userUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Name_User != "" {
		userUpdate.Name_User = data.Name_User
	}

	if data.Password != "" {
		hash_pass, errHash := HashPassword(data.Password)
		if errHash != nil {
			return -1, errHash
		}
		userUpdate.Password = hash_pass
	}

	if data.Foto_user != "" {
		userUpdate.Foto_user = data.Foto_user
	}

	if data.Address_user != "" {
		userUpdate.Address_user = data.Address_user
	}
	tx := repo.db.Save(&userUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataUser) InsertOwner(data user.Owner) (row int, err error) {
	owner := fromCoreOwner(data)

	tx := repo.db.Create(&owner)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataUser) SelectVerificationRequest() ([]user.UserCore, error) {
	var users []User
	// tx := repo.db.Where("user_owner = false AND foto_owner is null").Preload("Owner").Find(&users)

	tx := repo.db.Joins("inner join owners on owners.user_id = users.id").Where("user_owner= 0").Find(&users)
	// fmt.Println(users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	userCore := toCoreList(users)
	return userCore, nil
}

func (repo *dataUser) AdminApprove(data user.UserCore) (int, error) {
	var userUpdate User
	txDataOld := repo.db.First(&userUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Name_User != "" {
		userUpdate.Name_User = data.Name_User
	}

	if data.Password != "" {
		hash_pass, errHash := HashPassword(data.Password)
		if errHash != nil {
			return -1, errHash
		}
		userUpdate.Password = hash_pass
	}

	if data.Foto_user != "" {
		userUpdate.Foto_user = data.Foto_user
	}

	if data.Address_user != "" {
		userUpdate.Address_user = data.Address_user
	}

	if data.User_owner == false {
		userUpdate.User_owner = true
	}
	// fmt.Println(userUpdate.User_owner)
	tx := repo.db.Save(&userUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
