package storage

type UserServices struct {
	UserData map[int64]User
}

type User struct {
	UserID      int64
	Input       Action
	CurrServ    string
	ServiceName map[string]Service
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
			ServiceName: make(map[string]Service),
		}
		(*usersData).UserData[ID] = user
	}
}

func InitService(usersData *UserServices, ID int64, name string) {
	_, ok := (*usersData).UserData[ID].ServiceName[name]
	if !ok {
		service := Service{
			Login:    "",
			Password: "",
		}
		(*usersData).UserData[ID].ServiceName[name] = service
	}
}
