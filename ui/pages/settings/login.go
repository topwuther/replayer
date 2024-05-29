package settings

import (
	"github.com/go-resty/resty/v2"
)

type User struct {
	Username string
	Password string
	uuid     string
	perm     int
}

type Result struct {
	UUID string
	Perm int
}

type CheckForm struct {
	UUID string
}

func (user *User) Login() (bool, error) {
	form := user
	var loginResult Result
	server := SettingConfig.Server
	client := resty.New()
	_, err := client.R().SetBody(form).SetResult(&loginResult).Post("http://" + server + "/auth/login")
	if err != nil {
		return false, err
	}
	if loginResult.Perm != 0 {
		user.uuid = loginResult.UUID
		user.perm = loginResult.Perm
		return true, nil
	}
	return false, nil
}
