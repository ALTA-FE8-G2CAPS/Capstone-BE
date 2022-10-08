package usecase

import (
	"capstone-project/features/user"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockUserUsecase struct{}

// Test Logic Success
func (mock mockUserUsecase) LoginUser(data user.UserCore) (token string, err error) {
	return "token", nil
}

func (mock mockUserUsecase) InsertData(data user.UserCore) (row int, err error) {
	return 1, nil
}

func (mock mockUserUsecase) SelectAllUser() (data []user.UserCore, err error) {
	return data, nil
}

func (mock mockUserUsecase) SelectUserById(id int) (data user.UserCore, err error) {
	return user.UserCore{ID: 1, Name_User: "Jono", Email: "example@gmail.com", Password: "qwerty", Role: "user", Address_user: "Konohagakure", Foto_user: "bucketS3.aws.com", User_owner: false, Foto_owner: user.Owner{}, Created_At: time.Time{}, Updated_At: time.Time{}}, nil
}

func (mock mockUserUsecase) UpdateUser(data user.UserCore) (row int, err error) {
	return 1, nil
}

func (mock mockUserUsecase) DeleteUser(id int) (row int, err error) {
	return 1, nil
}

func (mock mockUserUsecase) InsertOwner(data user.Owner) (row int, err error) {
	return 1, nil
}

func (mock mockUserUsecase) SelectVerificationRequest() (data []user.UserCore, err error) {
	return []user.UserCore{}, nil
}

func (mock mockUserUsecase) AdminApprove(data user.UserCore) (row int, err error) {
	return 1, nil
}

//Test Logic Failed

type mockUserUsecaseFailed struct{}

func (mock mockUserUsecaseFailed) LoginUser(data user.UserCore) (token string, err error) {
	return "", fmt.Errorf("error login")
}

func (mock mockUserUsecaseFailed) InsertData(data user.UserCore) (row int, err error) {
	return 0, fmt.Errorf("error create user")
}

func (mock mockUserUsecaseFailed) SelectAllUser() (data []user.UserCore, err error) {
	return data, fmt.Errorf("error get all user")
}

func (mock mockUserUsecaseFailed) SelectUserById(id int) (data user.UserCore, err error) {
	return data, fmt.Errorf("error get user by id")
}

func (mock mockUserUsecaseFailed) UpdateUser(data user.UserCore) (row int, err error) {
	return 0, fmt.Errorf("error put data")
}

func (mock mockUserUsecaseFailed) DeleteUser(id int) (row int, err error) {
	return 0, fmt.Errorf("error delete user")
}

func (mock mockUserUsecaseFailed) InsertOwner(data user.Owner) (row int, err error) {
	return 0, fmt.Errorf("error post owner")
}

func (mock mockUserUsecaseFailed) SelectVerificationRequest() (data []user.UserCore, err error) {
	return data, fmt.Errorf("error get verification request")
}

func (mock mockUserUsecaseFailed) AdminApprove(data user.UserCore) (row int, err error) {
	return 0, fmt.Errorf("error admin approve")
}

func TestGetUserById(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		id := 1
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.GetUserById(id)
		assert.Nil(t, err)
		assert.Equal(t, "Jono", result.Name_User)
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		id := 10
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.GetUserById(id)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Name_User)
	})
}

func TestPostData(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		data := user.UserCore{Name_User: "Jono", Email: "example@gmail.com", Password: "qwerty", Address_user: "Konohagakure", Foto_user: "bucketS3.aws.com"}
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.PostData(data)
		assert.Nil(t, result, err)
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		data := user.UserCore{Name_User: "Jono", Email: "", Password: "qwerty", Address_user: "Konohagakure", Foto_user: "bucketS3.aws.com"}
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.PostData(data)
		assert.NotNil(t, result, err)
	})
	t.Run("Test Logic Failed when email is invalid", func(t *testing.T) {
		data := user.UserCore{Name_User: "Jono", Email: "example", Password: "qwerty", Address_user: "Konohagakure", Foto_user: "bucketS3.aws.com"}
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.PostData(data)
		assert.NotNil(t, result, err)
	})
}

func TestGetAllUser(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.GetAllUser()
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.GetAllUser()
		assert.NotNil(t, err)
		assert.Equal(t, -1, len(result))
	})
}

func TestPutData(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		data := user.UserCore{Name_User: "Jono", Email: "example@gmail.com", Password: "qwerty", Address_user: "Konohagakure", Foto_user: "bucketS3.aws.com"}
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.PutData(data)
		assert.Nil(t, result, err)
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		data := user.UserCore{Name_User: "Jono", Email: "example@gmail.com", Password: "qwerty", Address_user: "Konohagakure", Foto_user: "bucketS3.aws.com"}
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.PutData(data)
		assert.NotNil(t, result, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		id := 1
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.DeleteUser(id)
		assert.Nil(t, result, err)
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		id := 10
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.DeleteUser(id)
		assert.NotNil(t, result, err)
	})
}

func TestPostOwner(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		data := user.Owner{UserID: uint(1), Foto_owner: "bucketS3.aws.com"}
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.PostOwner(data)
		assert.Nil(t, result, err)
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		data := user.Owner{UserID: uint(1), Foto_owner: ""}
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.PostOwner(data)
		assert.NotNil(t, result, err)
	})
}

func TestPostLogin(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		data := user.UserCore{Email: "example@gmail.com", Password: "qwerty"}
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.PostLogin(data)
		assert.Nil(t, result, err)
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		data := user.UserCore{Email: "salah@gmail.com", Password: "salah"}
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.PostLogin(data)
		assert.NotNil(t, result, err)
	})
}

func TestGetVerificationRequest(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.GetVerificationRequest()
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.GetVerificationRequest()
		assert.NotNil(t, err)
		assert.Equal(t, -1, len(result))
	})
}

func TestAdminApprove(t *testing.T) {
	t.Run("Test Logic Success", func(t *testing.T) {
		var user user.UserCore
		user.ID = 1
		userUsecase := NewUserUsecase(mockUserUsecase{})
		result, err := userUsecase.AdminApprove(user)
		assert.Nil(t, result, err)
	})
	t.Run("Test Logic Failed", func(t *testing.T) {
		var user user.UserCore
		user.ID = 10
		userUsecase := NewUserUsecase(mockUserUsecaseFailed{})
		result, err := userUsecase.AdminApprove(user)
		assert.NotNil(t, result, err)
	})
}
