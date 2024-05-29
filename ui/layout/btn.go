package layout

import (
	"github.com/topwuther/replayer/ui/values"

	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Btn struct {
	Clickable *widget.Clickable
	Text      string
	Icon      *widget.Icon
}

func (btn *Btn) Layout(gtx C) D {
	if btn.Icon == nil {
		return material.Button(Theme, btn.Clickable, values.GetText(btn.Text)).Layout(gtx)
	}
	iconbutton := &IconButton{theme: Theme, icon: btn.Icon, word: values.GetText(btn.Text), button: btn.Clickable}
	return iconbutton.Layout(gtx)
}
