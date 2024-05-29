package dialog

import (
	giouilayout "gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/topwuther/replayer/ui"
	"github.com/topwuther/replayer/ui/layout"
	"github.com/topwuther/replayer/ui/values"
)

const PAGE_NAME = "dialog"

var (
	ckOK = new(widget.Clickable)
)

type Dialog struct {
	Title    string
	Text     string
	NextPage string
}

func (dialog *Dialog) Show(w *ui.Window) {
	w.Arguments = map[string]any{
		"title":    dialog.Title,
		"text":     dialog.Text,
		"nextpage": dialog.NextPage,
	}
	w.Page = PAGE_NAME
}

func Init(w *ui.Window) {

}

func Layout(w *ui.Window) layout.D {
	args := w.Arguments
	gtx := *w.GTX
	title := args["title"].(string)
	text := args["text"].(string)
	nextpage := args["nextpage"].(string)
	if ckOK.Clicked(gtx) {
		w.Page = nextpage
	}
	spacer := layout.Spacer{
		Axis:   giouilayout.Vertical,
		Weight: 5,
	}

	label := layout.Label{
		Text: text,
	}
	ch := func(gtx layout.C) layout.D {
		return giouilayout.Flex{Axis: giouilayout.Vertical}.Layout(gtx,
			giouilayout.Rigid(func(gtx layout.C) layout.D {
				return material.H3(layout.Theme, values.GetText(title)).Layout(gtx)
			}),
			giouilayout.Rigid(func(gtx layout.C) layout.D {
				return spacer.Layout(gtx)
			}),
			giouilayout.Rigid(func(gtx layout.C) layout.D {
				return label.Layout(gtx)
			}),
			giouilayout.Rigid(func(gtx layout.C) layout.D {
				return spacer.Layout(gtx)
			}),
			giouilayout.Rigid(func(gtx layout.C) layout.D {
				return giouilayout.Center.Layout(gtx, func(gtx layout.C) layout.D {
					return material.Button(layout.Theme, ckOK, values.GetText("ok")).Layout(gtx)
				})
			}),
		)
	}
	container := layout.Container{
		Width:    500,
		Height:   500,
		Children: ch,
	}
	return container.Layout(gtx)
}
