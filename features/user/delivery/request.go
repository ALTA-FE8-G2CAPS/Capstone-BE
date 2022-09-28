package delivery

import (
	"capstone-project/features/user"
)

type UserRequest struct {
	Nama_User    string `json:"nama_user" form:"nama_user"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Role         string `json:"role" form:"role"`
	Address_user string `json:"address_user" form:"address_user"`
	Foto_user    string `json:"foto_user" form:"foto_user"`
	User_owner   string `json:"user_owner" form:"user_owner"`
}

func ToCore(data UserRequest) user.UserCore {
	return user.UserCore{
		Nama_User:    data.Nama_User,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Address_user: data.Address_user,
		Foto_user:    data.Foto_user,
		User_owner:   data.User_owner,
	}
}
