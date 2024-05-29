package pages

import (
	"image/color"

	"github.com/topwuther/replayer/ui"
	"github.com/topwuther/replayer/ui/component"
	"github.com/topwuther/replayer/ui/layout"
	"github.com/topwuther/replayer/ui/pages/dialog"
	"github.com/topwuther/replayer/ui/pages/playlist"
	"github.com/topwuther/replayer/ui/pages/settings"

	giouilayout "gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

type (
	C = layout.C
	D = layout.D
)

type Page struct {
	Name     string
	Init     func(*ui.Window)
	Inited   bool
	Layout   func(*ui.Window) D
	ShowTabs bool
	LateInit bool
}

var (
pages = []Page{
	{
		Name:     playlist.PAGE_NAME,
		Init:     playlist.Init,
		Layout:   playlist.Layout,
		ShowTabs: true,
		LateInit: true,
	},
	{
		Name:     settings.PAGE_NAME,
		Init:     settings.Init,
		Layout:   settings.Layout,
		ShowTabs: true,
		LateInit: false,
	},
	{
		Name:     dialog.PAGE_NAME,
		Init:     dialog.Init,
		Layout:   dialog.Layout,
		ShowTabs: false,
		LateInit: false,
	},
}
	inited = false
	window *ui.Window
)

func PageHandler(w *ui.Window) {
	tabs := component.Tabs
	tabs.Callback = TabsHandler
	cktabs := make([]widget.Clickable, len(pages))
	w.Page = playlist.PAGE_NAME
	for {
		<-w.RenderStart
		window = w
		gtx := *w.GTX
		// Add all page to tab list
		tabs.TabList = []component.Tab{}
		for index, value := range pages {
			if !value.ShowTabs {
				continue
			}
			tabs.TabList = append(tabs.TabList, component.Tab{
				Title:     value.Name,
				Clickable: &cktabs[index],
			})
		}

		// Set theme
		darkTheme(w, w.DarkTheme)
		layout.Theme = w.Theme

		if !inited {
			for index, value := range pages {
				if value.LateInit {
					continue
				}
				value.Init(w)
				pages[index].Inited = true
			}
			inited = true
		}
		for index, value := range pages {
			// Match page
			if value.Name != w.Page {
				continue
			}

			// Init code
			if !value.Inited {
				value.Init(w)
				pages[index].Inited = true
			}

			// Set tabs
			if value.ShowTabs {
				sp := &layout.Spacer{
					Axis:   giouilayout.Vertical,
					Weight: 1,
				}
				giouilayout.Flex{Axis: giouilayout.Vertical}.Layout(gtx,
					giouilayout.Rigid(tabs.Layout),
					giouilayout.Rigid(sp.Layout),
					giouilayout.Rigid(func(gtx C) D {
						return value.Layout(w)
					}),
				)
				break
			}

			// Loop code
			value.Layout(w)
			break
		}
		w.RenderEnd <- struct{}{}
	}
}

func TabsHandler(selected int) {
	window.Page = pages[selected].Name
}

func blackBg(win *ui.Window) giouilayout.Dimensions {
	gtx := win.GTX
	c := win.Theme.Bg
	defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()
	defer win.Window.Invalidate()
	paint.ColorOp{Color: c}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return giouilayout.Dimensions{Size: gtx.Constraints.Max}
}

func darkTheme(win *ui.Window, darkmode bool) {
	switch darkmode {
	case true:
		blackBg(win)
		win.Theme.ContrastFg = color.NRGBA{R: 0xFF, G: 0x79, B: 0xC6, A: 0xff}
		win.Theme.ContrastBg = color.NRGBA{R: 0x4D, G: 0x4F, B: 0x68, A: 0xff}
		win.Theme.Fg = color.NRGBA{R: 0xFF, G: 0x79, B: 0xC6, A: 0xff}
		win.Theme.Bg = color.NRGBA{R: 0x28, G: 0x29, B: 0x36, A: 0xFF}
	case false:
		win.Theme.ContrastFg = color.NRGBA{R: 0xFF, G: 0xff, B: 0xff, A: 0xff}
		win.Theme.ContrastBg = color.NRGBA{R: 0x3f, G: 0x51, B: 0xb5, A: 0xff}
		win.Theme.Fg = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
		win.Theme.Bg = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	}
}
