package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	giouilayout "gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/topwuther/replayer/ui"
	"github.com/topwuther/replayer/ui/icon"
	"github.com/topwuther/replayer/ui/layout"
	"github.com/topwuther/replayer/ui/values"
	"gorm.io/gorm"
)

const PAGE_NAME = "settings"

type (
	C = giouilayout.Context
	D = giouilayout.Dimensions
)

type Config struct {
	Server   string
	User     User
	Language int
	Darkmode bool
	Repeat   bool
	Random   bool
	Single   bool
}

var (
	cfile         = "config.json"
	edServer      = new(widget.Editor)
	edUsername    = new(widget.Editor)
	edPassword    = new(widget.Editor)
	enumLanguage  = new(widget.Enum)
	enumDarkmode  = new(widget.Enum)
	blRepeat      = new(widget.Bool)
	blRandom      = new(widget.Bool)
	blSingle      = new(widget.Bool)
	ckSave        = new(widget.Clickable)
	Db            *gorm.DB
	SettingConfig Config
	msg           = layout.Label{}
)

func (c *Config) Show() {
	edServer.SetText(c.Server)
	edUsername.SetText(c.User.Username)
	edPassword.SetText(c.User.Password)
	enumLanguage.Value = strconv.Itoa(c.Language)
	if c.Darkmode {
		enumDarkmode.Value = strconv.Itoa(1)
	} else {
		enumDarkmode.Value = strconv.Itoa(0)
	}
	blRepeat.Value = c.Repeat
	blRandom.Value = c.Random
	blSingle.Value = c.Single
}

func (c *Config) Update(w *ui.Window) {
	values.SetUserLanguage(c.Language)
	w.DarkTheme = c.Darkmode
}

func Init(w *ui.Window) {
	fmt.Println("settings init")
	_, err := os.Stat(cfile)
	if os.IsNotExist(err) {
		return
	}

	// set config
	f, err := os.ReadFile(cfile)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(f, &SettingConfig); err != nil {
		panic(err)
	}
	SettingConfig.Show()
	SettingConfig.Update(w)
}

func Layout(w *ui.Window) D {
	gtx := *w.GTX
	form := layout.Form{
		Editor: []*layout.Editor{
			{
				Ed:    edServer,
				Hint:  "server",
				Width: 200,
			},
			{
				Ed:    edUsername,
				Hint:  "username",
				Width: 200,
			},
			{
				Ed:    edPassword,
				Hint:  "password",
				Width: 200,
			},
		},
		RadioBtnGrps: []*layout.RadioBtnGrp{
			{
				Enum:  enumLanguage,
				Title: "language",
				RadioBtns: []layout.RadioBtn{
					{
						Label: "english",
					},
					{
						Label: "chinese",
					},
				},
			},
			{
				Enum:  enumDarkmode,
				Title: "darkmode",
				RadioBtns: []layout.RadioBtn{
					{
						Label: "off",
					},
					{
						Label: "on",
					},
				},
			},
		},
		SwitchGrp: []*layout.Switch{
			{
				Label:    "repeat",
				Selected: blRepeat,
			},
			{
				Label:    "random",
				Selected: blRandom,
			},
			{
				Label:    "single",
				Selected: blSingle,
			},
		},
		BtnGrp: &layout.BtnGrp{
			Btns: []layout.Btn{
				{
					Clickable: ckSave,
					Text:      "save",
					Icon:      icon.SaveIcon,
				},
			},
		},
		Axis:   giouilayout.Vertical,
		Weight: 3,
	}
	if ckSave.Clicked(gtx) {
		language, _ := strconv.Atoi(enumLanguage.Value)
		darkmode, _ := strconv.Atoi(enumDarkmode.Value)
		user := User{
			Username: edUsername.Text(),
			Password: edPassword.Text(),
		}
		SettingConfig = Config{
			Server:   edServer.Text(),
			User:     user,
			Language: language,
			Darkmode: darkmode == 1,
			Repeat:   blRepeat.Value,
			Random:   blRandom.Value,
			Single:   blSingle.Value,
		}
		msg.Text = "save success"
		data, err := json.Marshal(SettingConfig)
		if err != nil {
			msg.Text = "save failed"
		}
		if err := os.WriteFile(cfile, data, 0644); err != nil {
			msg.Text = "save failed"
		}
		SettingConfig.Update(w)
		go func() {
			time.Sleep(1 * time.Second)
			msg.Text = ""
			w.Window.Invalidate()
		}()
	}
	return giouilayout.Flex{Axis: giouilayout.Vertical}.Layout(gtx,
		giouilayout.Rigid(func(gtx C) D {
			return giouilayout.Center.Layout(gtx, func(gtx C) D {
				return material.H1(layout.Theme, values.GetText("settings")).Layout(gtx)
			})
		}),
		giouilayout.Rigid(func(gtx C) D {
			return giouilayout.Center.Layout(gtx, func(gtx C) D {
				return form.Layout(gtx)
			})
		}),
		giouilayout.Rigid(func(gtx C) D {
			return giouilayout.Center.Layout(gtx, func(gtx C) D {
				return msg.Layout(gtx)
			})
		}),
	)
}
