package layout

import (
	"github.com/topwuther/replayer/ui/values"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Switch struct {
	Label    string
	Selected *widget.Bool
}

func (s *Switch) Layout(gtx C) D {
	return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			l := material.Body1(Theme, values.GetText(s.Label)+": ")
			return l.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Width: 5}.Layout),
		layout.Rigid(func(gtx C) D {
			return material.Switch(Theme, s.Selected, values.GetText(s.Label)).Layout(gtx)
		}))
}
