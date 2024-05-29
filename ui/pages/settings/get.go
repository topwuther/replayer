package settings

import "github.com/go-resty/resty/v2"

func (c *User) GetLoginStatus() bool {
	var loginResult Result
	form := CheckForm{
		UUID: c.uuid,
	}
	if c.perm == 0 {
		return false
	}
	server := SettingConfig.Server
	client := resty.New()
	_, err := client.R().SetBody(form).SetResult(&loginResult).Post("http://"+ server +"/auth/check")
	if err != nil {
		return false
	}
	if loginResult.Perm != 0 {
		return true
	}
	return false
}

func (c *User) GetUUID() string {
	return c.uuid
}
