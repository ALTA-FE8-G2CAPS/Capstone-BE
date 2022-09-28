package usecase

import (
	"capstone-project/features/user"
	"errors"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		data,
	}
}

func (usecase *userUsecase) PostData(data user.UserCore) (row int, err error) {
	if data.Nama_User == "" || data.Email == "" || data.Password == "" || data.Address_user == "" {
		return -1, errors.New("Data tidak boleh kosong")
	}
	row, err = usecase.userData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *userUsecase) PostLogin(data user.UserCore) (token string, err error) {
	if data.Email == "" || data.Password == "" {
		return "", errors.New("Email or Password is incorrect")
	}
	token, err = usecase.userData.LoginUser(data)
	if err != nil {
		return "", err
	}
	return token, err
}
