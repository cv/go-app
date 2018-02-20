package main

import (
	"fmt"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/app/drivers/mac"
)

func main() {
	app.DefaultLogger = app.NewConcurrentLogger(app.NewConsole(true))

	app.Run(&mac.Driver{
		MenubarConfig: mac.MenuBarConfig{
			OnPreference: func() {
				app.DefaultLogger.Log("Preferences clicked")
			},
		},
		DockURL: "menu",

		OnRun: func() {
			fmt.Println("OnRun")
			fmt.Println("app.Resources():", app.Resources())
			fmt.Println("app.Storage():", app.Storage())

			// testWindow(true)
			testWindow(false)
		},
		OnFocus: func() {
			fmt.Println("OnFocus")
		},
		OnBlur: func() {
			fmt.Println("OnBlur")
		},
		OnReopen: func(hasVisibleWindows bool) {
			fmt.Println("OnReopen hasVisibleWIndow:", hasVisibleWindows)
			if hasVisibleWindows {
				return
			}
			testWindow(false)
		},
		OnQuit: func() bool {
			fmt.Println("OnQuit")
			return true
		},
		OnExit: func() {
			fmt.Println("OnExit")
		},
	})
}

func testWindow(close bool) {
	win, _ := app.NewWindow(app.WindowConfig{
		Title:    "test window",
		X:        42,
		Y:        42,
		Width:    1024,
		MinWidth: 400,
		// MaxWidth:  300,
		Height:    600,
		MinHeight: 400,
		// NoResizable:     true,
		// NoClosable:      true,
		// NoMinimizable:   true,
		// TitlebarHidden:  true,
		BackgroundColor: "#282c34",
		Mac: app.MacWindowConfig{
			BackgroundVibrancy: app.VibeUltraDark,
		},
		DefaultURL: "webview",

		OnMove: func(x, y float64) {
			fmt.Printf("Window moved to x:%v y:%v\n", x, y)
		},
		OnResize: func(width, height float64) {
			fmt.Printf("Window resized to width:%v height:%v\n", width, height)
		},
		OnFocus: func() {
			fmt.Println("Window focused")
		},
		OnBlur: func() {
			fmt.Println("Window blured")
		},
		OnFullScreen: func() {
			fmt.Println("Window full screen")
		},
		OnExitFullScreen: func() {
			fmt.Println("Window exit full screen")
		},
		OnMinimize: func() {
			fmt.Println("Window minimized")
		},
		OnDeminimize: func() {
			fmt.Println("Window deminimized")
		},
		OnClose: func() bool {
			fmt.Println("Window close")
			return true
		},
	})

	fmt.Println("app name:", app.Name())

	x, y := win.Position()
	fmt.Printf("win.Positon() x:%v, x:%v\n", x, y)

	fmt.Printf("win.Move(x:%v, y: %v)\n", 42, 42)
	win.Move(42, 42)

	fmt.Println("win.Center()")
	win.Center()

	width, height := win.Size()
	fmt.Printf("win.Size() width:%v, height:%v\n", width, height)

	// fmt.Printf("win.Resize(x:%v, y: %v)\n", 1340, 720)
	// win.Resize(1340, 720)

	// fmt.Println("win.ToggleFullScreen()")
	// win.ToggleFullScreen()

	// go func() {
	// 	fmt.Println("win.ToggleFullScreen()")
	// 	time.Sleep(2 * time.Second)
	// 	win.ToggleFullScreen()
	// }()

	// fmt.Println("win.ToggleMinimize()")
	// win.ToggleMinimize()

	// fmt.Println("win.ToggleMinimize()")
	// win.ToggleMinimize()

	// win.Focus()

	if close {
		fmt.Println("win.Close()")
		win.Close()
	}

	fmt.Println("Window tests OK")
}
