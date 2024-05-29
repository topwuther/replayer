package component

import (
	"image"

	"github.com/topwuther/replayer/ui/layout"
	"github.com/topwuther/replayer/ui/values"

	giouilayout "gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var Tabs = &tabs{}

type Tab struct {
	Clickable *widget.Clickable
	Title     string
}

type tabs struct {
	list     giouilayout.List
	TabList  []Tab
	selected int
	Callback func(selected int)
}

func (tabs *tabs) Layout(gtx layout.C) layout.D {
	return tabs.list.Layout(gtx, len(tabs.TabList), func(gtx layout.C, tabIdx int) layout.D {
		t := &tabs.TabList[tabIdx]
		if t.Clickable.Clicked(gtx) {
			tabs.selected = tabIdx
			tabs.Callback(tabs.selected)
		}
		var tabWidth int
		return giouilayout.Stack{Alignment: giouilayout.S}.Layout(gtx,
			giouilayout.Stacked(func(gtx layout.C) layout.D {
				dims := material.Clickable(gtx, t.Clickable, func(gtx layout.C) layout.D {
					return giouilayout.UniformInset(unit.Dp(12)).Layout(gtx,
						material.H6(layout.Theme, values.GetText(t.Title)).Layout,
					)
				})
				tabWidth = dims.Size.X
				return dims
			}),
			giouilayout.Stacked(func(gtx layout.C) layout.D {
				if tabs.selected != tabIdx {
					return giouilayout.Dimensions{}
				}
				tabHeight := gtx.Dp(unit.Dp(4))
				tabRect := image.Rect(0, 0, tabWidth, tabHeight)
				paint.FillShape(gtx.Ops, layout.Theme.Palette.ContrastBg, clip.Rect(tabRect).Op())
				return giouilayout.Dimensions{
					Size: image.Point{X: tabWidth, Y: tabHeight},
				}
			}),
		)
	})
}

func (tabs *tabs) Update(page string) {
	for i,tab := range tabs.TabList {
		if tab.Title == page {
			tabs.selected = i
		}
	}
}
