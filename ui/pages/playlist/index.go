package playlist

import (
	_ "embed"
	"fmt"
	"io"
	"time"

	giouilayout "gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	player1 "github.com/topwuther/replayer/libreplayer/player"
	"github.com/topwuther/replayer/ui"
	"github.com/topwuther/replayer/ui/icon"
	"github.com/topwuther/replayer/ui/layout"
	"github.com/topwuther/replayer/ui/pages/dialog"
	"github.com/topwuther/replayer/ui/pages/settings"
	"github.com/topwuther/replayer/ui/values"
)

const PAGE_NAME = "playlist"

type (
	C = giouilayout.Context
	D = giouilayout.Dimensions
)

var (
	player = player1.NewPlayer()
	// record = librecord.NewRecord("a.wav")

	grid      = new(component.GridState)
	tl        = &TableLayout{}
	playIndex = 0
	cklist    = make([]widget.Clickable, 10)
	ckplay    = new(widget.Clickable)
	cknext    = new(widget.Clickable)
	ckprev    = new(widget.Clickable)
	progress  = new(widget.Float)
)

func Init(w *ui.Window) {
	// Init audio device
	// record.SetSourceByName("WordForum USB Mono")
	// if record.Err != nil {
	// 	fmt.Println("record message:", record.Err.Error())
	// }
	// player.SetSinkByName("Built-in Audio Analog Stereo")
	// if player.Err != nil {
	// 	fmt.Println("player message:", player.Err.Error())
	// }

	// Set progressbar
	go func() {
		for {
			time.Sleep(time.Second)
			if player.Downloading {
				continue
			}
			progress.Value = player.Progress()
			if progress.Value > 0.999 {
				autoAction()
				if player.Err != nil {
					dialog := dialog.Dialog{
						Title:    "play failed",
						Text:     player.Err.Error(),
						NextPage: PAGE_NAME,
					}
					dialog.Show(w)
				}
			}
			w.Window.Invalidate()
		}
	}()
	// record.Start()
}

func Layout(w *ui.Window) D {
	gtx := *w.GTX
	if len(musics) == 0 {
		tl.Build(w)
	}
	t := &layout.Table{
		Grid:   grid,
		Height: 40,
		Rows: []layout.Column{
			{
				Title:  "index",
				Weight: 1,
				Layout: tl.Index,
			},
			{
				Title:  "name",
				Weight: 8,
				Layout: tl.Name,
			},
			{
				Title:  "singer",
				Weight: 8,
				Layout: tl.Singer,
			},
			{
				Title:  "play",
				Weight: 1,
				Layout: tl.PlayBtn,
			},
		},
	}
	playbtn := layout.Btn{
		Clickable: ckplay,
	}
	switch player.IsPlaying() {
	case true:
		playbtn.Text = "pause"
		playbtn.Icon = icon.PauseIcon
	case false:
		playbtn.Text = "play"
		playbtn.Icon = icon.PlayIcon
	}
	mediaBtns := layout.BtnGrp{
		Btns: []layout.Btn{
			{
				Text:      "previous",
				Icon:      icon.PreviousIcon,
				Clickable: ckprev,
			},
			playbtn,
			{
				Text:      "next",
				Icon:      icon.NextIcon,
				Clickable: cknext,
			},
		},
	}
	// action handler
	switch {
	case ckplay.Clicked(gtx):
		switch player.IsPlaying() {
		case true:
			player.Pause()
		case false:
			player.Play(&musics[playIndex])
		}
	case progress.Update(gtx):
		player.Seek(progress.Value, io.SeekStart)
	case ckprev.Clicked(gtx):
		prevAction()
		if player.Err != nil {
			dialog := dialog.Dialog{
				Title:    "play failed",
				Text:     player.Err.Error(),
				NextPage: PAGE_NAME,
			}
			dialog.Show(w)
		}
	case cknext.Clicked(gtx):
		nextAction()
		if player.Err != nil {
			dialog := dialog.Dialog{
				Title:    "play failed",
				Text:     player.Err.Error(),
				NextPage: PAGE_NAME,
			}
			dialog.Show(w)
		}
	}

	for i := range cklist {
		if cklist[i].Clicked(gtx) {
			player.Stop()
			player.Play(&musics[i])
			if player.Err != nil {
				dialog := dialog.Dialog{
					Title:    "play failed",
					Text:     player.Err.Error(),
					NextPage: settings.PAGE_NAME,
				}
				fmt.Println(player.Err.Error())
				dialog.Show(w)
			}
			playIndex = i
		}
	}

	return giouilayout.Flex{Axis: giouilayout.Vertical}.Layout(gtx,
		giouilayout.Rigid(func(gtx C) D {
			return giouilayout.Center.Layout(gtx, func(gtx C) D {
				return material.H1(layout.Theme, values.GetText("playlist")).Layout(gtx)
			})
		}),
		giouilayout.Rigid(func(gtx C) D {
			ch := func(gtx layout.C) layout.D {
				return giouilayout.Flex{}.Layout(gtx,
					giouilayout.Rigid(t.Layout),
					giouilayout.Rigid(giouilayout.Spacer{Height: unit.Dp(gtx.Constraints.Max.Y)}.Layout))
			}
			container := layout.Container{
				Width:    int(float32(gtx.Constraints.Max.X) * 0.8),
				Height:   int(float32(gtx.Constraints.Max.Y) * 0.8),
				Children: ch,
			}
			return container.Layout(gtx)
		}),
		giouilayout.Rigid(func(gtx C) D {
			return giouilayout.Flex{Axis: giouilayout.Horizontal}.Layout(gtx,
				giouilayout.Rigid(func(gtx C) D {
					spacer := layout.Spacer{
						Axis:   giouilayout.Horizontal,
						Weight: 2,
					}
					return spacer.Layout(gtx)
				}),
				giouilayout.Rigid(func(gtx C) D {
					return mediaBtns.Layout(gtx)
				}),
				giouilayout.Rigid(func(gtx C) D {
					spacer := layout.Spacer{
						Axis:   giouilayout.Horizontal,
						Weight: 2,
					}
					return spacer.Layout(gtx)
				}),
				giouilayout.Rigid(func(gtx C) D {
					gtx.Constraints.Min.X = int(float32(gtx.Constraints.Max.X))
					return giouilayout.Flex{Axis: giouilayout.Vertical}.Layout(gtx,
						giouilayout.Rigid(func(gtx giouilayout.Context) giouilayout.Dimensions {
							title := player.Title()
							return material.Body1(layout.Theme, title).Layout(gtx)
						}),
						giouilayout.Rigid(func(gtx C) D {
							return material.Slider(layout.Theme, progress).Layout(gtx)
						}))
				}),
			)
		}),
	)
}
