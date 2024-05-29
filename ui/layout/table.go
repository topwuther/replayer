package layout

import (
	"image"
	"image/color"

	"github.com/topwuther/replayer/ui/values"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type Column struct {
	Title  string
	Weight int
	Layout []func(C) D
}

type Table struct {
	Rows   []Column
	Grid   *component.GridState
	Height int
}

func (table *Table) Layout(gtx C) D {
	border := widget.Border{
		Color: color.NRGBA{A: 255},
		Width: unit.Dp(1),
	}
	headingLabel := material.Body1(Theme, "")
	headingLabel.Alignment = text.Middle
	headingLabel.MaxLines = 1

	gtx.Constraints.Min = image.Point{}

	inset := layout.UniformInset(unit.Dp(2))
	if table.Height == 0 {
		table.Height = inset.Layout(gtx, headingLabel.Layout).Size.Y
	}

	col := 0
	width := 0
	if len(table.Rows) != 0 {
		col = len(table.Rows[0].Layout)
	}
	for _, c := range table.Rows {
		width += c.Weight
	}

	return component.Table(Theme, table.Grid).Layout(gtx, col, len(table.Rows),
		func(axis layout.Axis, index, constraint int) int {
			widthUnit := int(constraint / width * table.Rows[index%len(table.Rows)].Weight)
			switch axis {
			case layout.Horizontal:
				return int(widthUnit)
			default:
				return table.Height
			}
		},
		func(gtx C, col int) D {
			return border.Layout(gtx, func(gtx C) D {
				return inset.Layout(gtx, func(gtx C) D {
					headingLabel.Text = values.GetText(table.Rows[col].Title)
					return headingLabel.Layout(gtx)
				})
			})
		},
		func(gtx C, row, col int) D {
			return table.Rows[col].Layout[row](gtx)
		},
	)
}
