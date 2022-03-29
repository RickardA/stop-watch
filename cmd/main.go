package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var white = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

type StopWatch struct {
	timer *time.Ticker
}

func main() {
	sw := StopWatch{}

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	content := container.New(layout.NewGridLayout(6))

	for i := 1; i <= 6; i++ {
		content.Add(createLaneContainer(i))
	}

	myWindow.Resize(fyne.NewSize(800, 300))

	timerText, topBarContainer := createTopBarContainer()

	bottomBarContainer := createBottomBarContainer(func() {
		// Start
		sw.timer = time.NewTicker(time.Millisecond)

		go func(timerText *canvas.Text) {
			counter := 0
			for range sw.timer.C {
				counter += 1
				timerText.Text = msToTime(counter)
				timerText.Refresh()
			}
		}(timerText)
	}, func() {
		sw.timer.Stop()
	})

	mainContainer := container.New(layout.NewBorderLayout(topBarContainer, bottomBarContainer, nil, nil),
		topBarContainer, bottomBarContainer, content)

	myWindow.SetContent(mainContainer)
	myWindow.Show()
	myApp.Run()
}

func msToTime(ms int) string {
	minutes := 0
	seconds := 0
	centiSeconds := 0

	centiMs := ms % 1000
	if centiMs != 0 {
		centiSeconds = centiMs / 10
	}

	ms = ms - centiMs

	tempSec := ms / 1000

	seconds = tempSec % 60

	tempSec = tempSec - seconds

	minutes = tempSec / 60

	return fmt.Sprintf("%02d:%02d:%02d", minutes, seconds, centiSeconds)
}

func createTopBarContainer() (*canvas.Text, *fyne.Container) {
	timerText := canvas.NewText("00:00:00", color.White)
	timerText.TextSize = 50
	return timerText, container.New(layout.NewHBoxLayout(), layout.NewSpacer(), timerText, layout.NewSpacer())
}

func createBottomBarContainer(startCallback func(), stopCallback func()) *fyne.Container {
	startBtn := widget.NewButton("STARTA", startCallback)
	stopBtn := widget.NewButton("STANNA", stopCallback)
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
