package layout

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Terminal struct {
	Title  string
	Ed     *widget.Editor
	Text   string
	Width  int
	Height int
}

func (terminal *Terminal) Layout(gtx C) D {
	title := Label{Text: terminal.Title}
	text := material.Editor(Theme, terminal.Ed, "")
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(title.Layout),
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Max.Y = terminal.Height
			gtx.Constraints.Max.X = terminal.Width
			gtx.Constraints.Min.Y = terminal.Height
			gtx.Constraints.Min.X = terminal.Width
			return layout.Background{}.Layout(gtx,
				func(gtx C) D {
					defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
					paint.Fill(gtx.Ops, color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
					return D{Size: gtx.Constraints.Min}
				}, text.Layout)
		}),
	)
}
