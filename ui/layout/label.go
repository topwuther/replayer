package layout

import (
	"github.com/topwuther/replayer/ui/values"

	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Label struct {
	Text string
	Size int
	AppendText string
}

func (label *Label) Layout(gtx C) D {
	if label.Size == 0 {
		label.Size = 20
	}
	return material.Label(Theme, unit.Sp(label.Size), values.GetText(label.Text)+label.AppendText).Layout(gtx)
}
