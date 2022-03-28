package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var white = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	content := container.New(layout.NewGridLayout(6))

	for i := 1; i <= 6; i++ {
		content.Add(createLaneContainer(i))
	}

	myWindow.Resize(fyne.NewSize(800, 300))

	topBarContainer := createTopBarContainer()

	bottomBarContainer := createBottomBarContainer()

	mainContainer := container.New(layout.NewBorderLayout(topBarContainer, bottomBarContainer, nil, nil),
		topBarContainer, bottomBarContainer, content)

	myWindow.SetContent(mainContainer)
	myWindow.Show()
	myApp.Run()
}

func createTopBarContainer() *fyne.Container {
	timerText := canvas.NewText("00:00:00", color.White)
	timerText.TextSize = 50
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), timerText, layout.NewSpacer())
}

func createBottomBarContainer() *fyne.Container {
	startBtn := widget.NewButton("STARTA", func() {})
	stopBtn := widget.NewButton("STANNA", func() {})
	btnContainer := container.New(layout.NewHBoxLayout(), startBtn, stopBtn)
	bottomBarContainer := container.New(layout.NewCenterLayout(), layout.NewSpacer(), btnContainer, layout.NewSpacer())

	return bottomBarContainer
}

func createLaneContainer(laneNumber int) *fyne.Container {
	laneTitle := canvas.NewText(fmt.Sprintf("Bana %d", laneNumber), white)
	timeText := canvas.NewText("00.00.00", white)
	laneTitle.TextSize = 20
	timeText.TextSize = 20
	vbox := container.New(layout.NewVBoxLayout(), laneTitle, timeText)
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), vbox, layout.NewSpacer())
}
