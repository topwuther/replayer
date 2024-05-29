package layout

import (
	"github.com/topwuther/replayer/ui/values"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Editor struct {
	Ed      *widget.Editor
	Hint    string
	Width   int
	Height  int
	Disable bool
}

func (ed *Editor) Layout(gtx C) D {
	ed.Ed.ReadOnly = ed.Disable
	e := material.Editor(Theme, ed.Ed, values.GetText(ed.Hint))
	e.Font.Style = font.Italic
	border := widget.Border{Color: Theme.Fg, CornerRadius: unit.Dp(8), Width: unit.Dp(2)}

	return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			l := material.Body1(Theme, values.GetText(ed.Hint)+": ")
			return l.Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Width: 5}.Layout),
		layout.Rigid(func(gtx C) D {
			return border.Layout(gtx, func(gtx C) D {
				if ed.Width != 0 {
					gtx.Constraints.Min.X = ed.Width
					gtx.Constraints.Max.X = ed.Width
				}
				if ed.Height != 0 {
					gtx.Constraints.Min.Y = ed.Width
					gtx.Constraints.Max.Y = ed.Width
				}
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
			})
		}))
}
