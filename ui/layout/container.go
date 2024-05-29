package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

type Container struct {
	Width    int
	Height   int
	Children func(gtx C) D
}

func (container *Container) Layout(gtx layout.Context) D {
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Max.X = gtx.Dp(unit.Dp(container.Width))
		gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(container.Height))
		return container.Children(gtx)
	})
}
