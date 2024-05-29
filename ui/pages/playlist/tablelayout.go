package playlist

import (
	"strconv"

	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/go-resty/resty/v2"
	player1 "github.com/topwuther/replayer/libreplayer/player"
	"github.com/topwuther/replayer/ui"
	"github.com/topwuther/replayer/ui/icon"
	"github.com/topwuther/replayer/ui/layout"
	"github.com/topwuther/replayer/ui/pages/dialog"
	"github.com/topwuther/replayer/ui/pages/settings"
)

type TableLayout struct {
	Index    []func(C) D
	Filename []string
	Name     []func(C) D
	Singer   []func(C) D
	PlayBtn  []func(C) D
}

type GetForm struct {
	Id   int
	UUID string
}

var musics = make([]player1.Music, 0)

func (t *TableLayout) Build(w *ui.Window) {
	t.Index = []func(C) D{}
	t.Name = []func(C) D{}
	t.Singer = []func(C) D{}

	user := settings.SettingConfig.User
	if !user.GetLoginStatus() {
		res, err := user.Login()
		switch {
		case err != nil:
			dialog := dialog.Dialog{
				Title:    "login failed",
				Text:     err.Error(),
				NextPage: settings.PAGE_NAME,
			}
			dialog.Show(w)
		case !res:
			dialog := dialog.Dialog{
				Title:    "login failed",
				Text:     "username or password has wrong",
				NextPage: settings.PAGE_NAME,
			}
			dialog.Show(w)
		}
	}
	payload := GetForm{
		Id:   0,
		UUID: user.GetUUID(),
	}
	server := settings.SettingConfig.Server
	client := resty.New()
	_, err := client.R().SetBody(payload).SetResult(&musics).Post("http://" + server + "/music/query")
	if err != nil {
		dialog := dialog.Dialog{
			Title:    "query music list failed",
			Text:     err.Error(),
			NextPage: settings.PAGE_NAME,
		}
		dialog.Show(w)
	}
	for index, value := range musics {
		musics[index].UUID = payload.UUID
		t.Index = append(t.Index, func(gtx C) D {
			bt := material.Body1(layout.Theme, strconv.Itoa(index+1))
			bt.Alignment = text.Middle
			bt.MaxLines = 1
			return bt.Layout(gtx)
		})
		t.Name = append(t.Name, func(gtx C) D {
			bt := material.Body1(layout.Theme, value.Name)
			bt.Alignment = text.Middle
			bt.MaxLines = 1
			return bt.Layout(gtx)
		})
		t.Singer = append(t.Singer, func(gtx C) D {
			bt := material.Body1(layout.Theme, value.Singer)
			bt.Alignment = text.Middle
			bt.MaxLines = 1
			return bt.Layout(gtx)
		})
		t.Filename = append(t.Filename, value.Filename)
		t.PlayBtn = append(t.PlayBtn, func(gtx C) D {
			btn := &layout.Btn{
				Text:      "play",
				Clickable: &cklist[index],
				Icon:      icon.PlayIcon,
			}
			if gtx.Constraints.Min.X < 75 {
				btn.Text = ""
			}
			return btn.Layout(gtx)
		})
	}

}
