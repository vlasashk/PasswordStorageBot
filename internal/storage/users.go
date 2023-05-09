package storage

import "gorm.io/gorm"

type UserServices struct {
	UserData map[int64]User
}

type User struct {
	UserID      int64
	Input       Action
	CurrServ    string
	ServiceName Service
}

type Service struct {
	Login    string
	Password string
}

type Action struct {
	Cmd   int //0 - command, 1 - set, 2 - get, 3 - del
	Login bool
	Pass  bool
}

type UserStorage struct {
	gorm.Model
	UserID   int64            `gorm:"primaryKey"`
	Services []ServiceStorage `gorm:"ForeignKey:UserID"`
}

type ServiceStorage struct {
	gorm.Model
	Name     string `gorm:"primaryKey"`
	UserID   int64
	Login    string
	Password string
}

func InitUsersStorage() *UserServices {
	usersData := &UserServices{
		UserData: make(map[int64]User),
	}
	return usersData
}

func InitUser(usersData *UserServices, ID int64) {
	_, ok := (*usersData).UserData[ID]
	if !ok {
		user := User{
			UserID:      ID,
			Input:       Action{Cmd: 0, Login: false, Pass: false},
			CurrServ:    "",
			ServiceName: Service{"", ""},
		}
		(*usersData).UserData[ID] = user
	}
}
