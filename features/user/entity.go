package user

import "time"

type UserCore struct {
	ID           uint
	Nama_User    string
	Email        string
	Password     string
	Role         string
	Address_user string
	Foto_user    string
	User_owner   string
	Foto_owner   Owner
	Created_At   time.Time
	Updated_At   time.Time
}

type Owner struct {
	Foto_owner string
}

type UsecaseInterface interface {
	GetAllUser() (data []UserCore, err error)
	GetUserById(id int) (data UserCore, err error)
	PostLogin(data UserCore) (token string, err error)
	PostData(data UserCore) (row int, err error)
	PutData(data UserCore) (row int, err error)
	DeleteUser(id int) (row int, err error)
	PostOwner(data Owner) (row int, err error)
}

type DataInterface interface {
	SelectAllUser() (data []UserCore, err error)
	SelectUserById(id int) (data UserCore, err error)
	InsertData(data UserCore) (row int, err error)
	LoginUser(data UserCore) (token string, err error)
	UpdateUser(data UserCore) (row int, err error)
	DeleteUser(id int) (row int, err error)
	InsertOwner(data Owner) (row int, err error)
}
