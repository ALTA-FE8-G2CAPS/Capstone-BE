package delivery

import (
	"capstone-project/features/user"
)

type UserRequest struct {
	Name_User    string `json:"name_user" form:"name_user"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Role         string `json:"role" form:"role"`
	Address_user string `json:"address_user" form:"address_user"`
	Foto_user    string `json:"foto_user" form:"foto_user"`
	User_owner   bool   `json:"user_owner" form:"user_owner"`
}

type OwnerRequest struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	Foto_owner string `json:"foto_owner" form:"foto_owner"`
}

func ToCoreOwner(data OwnerRequest) user.Owner {
	return user.Owner{
		UserID:     data.UserID,
		Foto_owner: data.Foto_owner,
	}
}

func ToCore(data UserRequest) user.UserCore {
	return user.UserCore{
		Name_User:    data.Name_User,
		Email:        data.Email,
		Password:     data.Password,
		Role:         data.Role,
		Address_user: data.Address_user,
		Foto_user:    data.Foto_user,
		User_owner:   data.User_owner,
	}
}
