package ui

import (
	"github.com/topwuther/replayer/ui/component"
	"github.com/topwuther/replayer/ui/values"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

type Window struct {
	Window      *app.Window
	Page        string
	Arguments   map[string]any
	Theme       *material.Theme
	RenderStart chan struct{}
	RenderEnd   chan struct{}
	GTX         *layout.Context
	Quit        chan struct{}
	DarkTheme   bool
}

func CreateWindow(title string) *Window {
	// Resolution
	giouiSize := app.Size(1920, 1080)
	appMinSize := app.MinSize(1280, 720)

	appTitle := app.Title(values.GetText(title))
	giouiWindow := new(app.Window)
	giouiWindow.Option(giouiSize, appMinSize, appTitle)
	win := &Window{
		Window:      giouiWindow,
		Page:        "/",
		Arguments:   make(map[string]any),
		Theme:       material.NewTheme(),
		RenderStart: make(chan struct{}),
		RenderEnd:   make(chan struct{}),
		Quit:        make(chan struct{}),
		DarkTheme:   false,
	}
	go EventHandler(win)
	return win
}

func EventHandler(win *Window) {
	var ops op.Ops
	for {
		switch event := win.Window.Event().(type) {
		case app.DestroyEvent:
			win.Quit <- struct{}{}
		case app.FrameEvent:
			gtx := app.NewContext(&ops, event)
			win.GTX = &gtx
			win.RenderStart <- struct{}{}
			component.Tabs.Update(win.Page)
			<-win.RenderEnd
			event.Frame(gtx.Ops)
		}
	}
}
