package main

import (
	"github.com/Waitfish/GrokDebugger/src/tabs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.NewWithID("EatFish")
	// icon := app.New().Icon()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("Grok Debugger by Waitfish")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Init config

	mainCotainer := container.NewAppTabs(
		tabs.RunIndexTab(myApp, myWindow),
	)
	myWindow.SetContent(mainCotainer)
	myWindow.ShowAndRun()
}
