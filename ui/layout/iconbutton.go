package layout

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type IconButton struct {
	theme  *material.Theme
	button *widget.Clickable
	icon   *widget.Icon
	word   string
}

func (b *IconButton) Layout(gtx layout.Context) layout.Dimensions {
	return material.ButtonLayout(b.theme, b.button).Layout(gtx, func(gtx C) D {
		return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
			iconAndLabel := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}

			layIcon := layout.Rigid(func(gtx C) D {
				return layout.Inset{}.Layout(gtx, func(gtx C) D {
					var d D
					if b.icon != nil {
						size := gtx.Dp(unit.Dp(60)) - 2*gtx.Dp(unit.Dp(18))
						gtx.Constraints = layout.Exact(image.Pt(size, size))
						d = b.icon.Layout(gtx, b.theme.ContrastFg)
					}
					return d
				})
			})

			layLabel := layout.Rigid(func(gtx C) D {
				return layout.Inset{}.Layout(gtx, func(gtx C) D {
					l := material.Body1(b.theme, b.word)
					l.Color = b.theme.Palette.ContrastFg
					return l.Layout(gtx)
				})
			})

			return iconAndLabel.Layout(gtx, layIcon, layLabel)
		})
	})
}
