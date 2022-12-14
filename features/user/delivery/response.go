package delivery

import (
	"capstone-project/features/user"
)

type UserResponse struct {
	ID           uint   `json:"id" form:"id"`
	Name_User    string `json:"name_user" form:"name_user"`
	Email        string `json:"email" form:"email"`
	Role         string `json:"role" form:"role"`
	Address_user string `json:"address_user" form:"address_user"`
	Foto_user    string `json:"foto_user" form:"foto_user"`
	User_owner   bool   `json:"user_owner" form:"user_owner"`
}

type OwnerResponse struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	Foto_owner string `json:"foto_owner" form:"foto_owner"`
}

func FromCoreOwner(data user.Owner) OwnerResponse {
	return OwnerResponse{
		UserID:     data.UserID,
		Foto_owner: data.Foto_owner,
	}
}

func FromCore(data user.UserCore) UserResponse {
	return UserResponse{
		ID:           data.ID,
		Name_User:    data.Name_User,
		Email:        data.Email,
		Role:         data.Role,
		Address_user: data.Address_user,
		Foto_user:    data.Foto_user,
		User_owner:   data.User_owner,
	}

}

func FromCoreList(data []user.UserCore) []UserResponse {
	var list []UserResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
