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
	if data.Name_User == "" || data.Email == "" || data.Password == "" || data.Address_user == "" {
		return -1, errors.New("data tidak boleh kosong")
	}
	row, err = usecase.userData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *userUsecase) PostLogin(data user.UserCore) (token string, err error) {
	if data.Email == "" || data.Password == "" {
		return "", errors.New("email or password is incorrect")
	}
	token, err = usecase.userData.LoginUser(data)
	if err != nil {
		return "", err
	}
	return token, err
}

func (usecase *userUsecase) GetAllUser() (data []user.UserCore, err error) {
	data, err = usecase.userData.SelectAllUser()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (usecase *userUsecase) GetUserById(id int) (data user.UserCore, err error) {
	data, err = usecase.userData.SelectUserById(id)
	if err != nil {
		return user.UserCore{}, err
	} else if data.ID == 0 {
		return user.UserCore{}, errors.New("user not found")
	} else { // data.ID != 0
		return data, err
	}
}

func (usecase *userUsecase) PutData(data user.UserCore) (row int, err error) {

	row, err = usecase.userData.UpdateUser(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *userUsecase) DeleteUser(id int) (row int, err error) {
	row, err = usecase.userData.DeleteUser(id)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (usecase *userUsecase) PostOwner(data user.Owner) (row int, err error) {
	row, err = usecase.userData.InsertOwner(data)
	if err != nil {
		return -1, err
	}
	return row, err
}
