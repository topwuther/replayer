package settings

import "github.com/go-resty/resty/v2"

type LogoutForm struct {
	UUID string
}

func (user *User) Logout() bool {
	var res Result
	payload := LogoutForm {
		UUID: user.uuid,
	}
	server := SettingConfig.Server
	client := resty.New()
	_, err := client.R().SetBody(payload).SetResult(&res).Post("http://"+ server +"/auth/login")
	if err != nil {
		return false
	}
	user.perm = res.Perm
	return user.perm == 0
}
