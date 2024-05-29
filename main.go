package main

import (
	// _ "net/http/pprof"

	// "net/http"
	"os"

	"gioui.org/app"
	"github.com/topwuther/replayer/ui"
	"github.com/topwuther/replayer/ui/pages"
)

func main() {
	// go func(){
	// 	http.ListenAndServe("localhost:8080", nil)
	// }()
	win := ui.CreateWindow("ktv")
	go pages.PageHandler(win)
	go func() {
		for {
			<-win.Quit
			os.Exit(0)
		}
	}()
	app.Main()
}
