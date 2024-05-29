package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

type Spacer struct {
	Axis   layout.Axis
	Weight float32
}

func (spacer *Spacer) Layout(gtx C) D {
	res := layout.Spacer{}
	switch spacer.Axis {
	case layout.Horizontal:
		res.Width = unit.Dp(float32(gtx.Constraints.Max.X) / 100 * spacer.Weight)
	case layout.Vertical:
		res.Height = unit.Dp(float32(gtx.Constraints.Max.Y) / 100 * spacer.Weight)
	}
	return res.Layout(gtx)
}
