package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/RickardA/stop-watch/internal/pkg/stop_watch"
)

var white = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

type App struct {
	mainWindow    *fyne.Window
	mainTimerText *canvas.Text
	app           *fyne.App
	Sw            *stop_watch.StopWatch
	lanes         []*Lane
	stopChan      *chan int
}

func NewApp(name string, stopChan *chan int) {
	myApp := app.New()
	mainWindow := myApp.NewWindow(name)

	app := App{
		mainWindow: &mainWindow,
		app:        &myApp,
		Sw:         stop_watch.NewStopWatch(),
		stopChan:   stopChan,
	}

	mainWindow.Resize(fyne.NewSize(800, 300))

	mainTimerText, mainTimerContainer := app.createMainTimer()
	app.mainTimerText = mainTimerText

	mainControlsContainer := app.createMainControls()

	lanes := app.createLanes()

	mainContainer := container.New(layout.NewBorderLayout(mainTimerContainer, mainControlsContainer, nil, nil),
		mainTimerContainer, mainControlsContainer, lanes)

	mainWindow.SetContent(mainContainer)
	mainWindow.Show()

	go app.listenOnStopChan()

	// Will not continue after this....
	myApp.Run()
}

func (a *App) listenOnStopChan() {
	for laneNum := range *a.stopChan {
		a.lanes[laneNum-1].text.Text = a.Sw.GetCurrentCountFormatted()
		a.lanes[laneNum-1].text.Refresh()
	}
}

func (a *App) createMainTimer() (*canvas.Text, *fyne.Container) {
	timerText := canvas.NewText("00:00:00", color.White)
	timerText.TextSize = 50
	container := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), timerText, layout.NewSpacer())
	return timerText, container
}

func (a *App) createMainControls() *fyne.Container {
	startBtn := widget.NewButton("STARTA", a.startBtnPressed)
	stopBtn := widget.NewButton("STANNA", a.stopBtnPressed)
	btnContainer := container.New(layout.NewHBoxLayout(), startBtn, stopBtn)
	bottomBarContainer := container.New(layout.NewCenterLayout(), layout.NewSpacer(), btnContainer, layout.NewSpacer())

	return bottomBarContainer
}

func (a *App) createLanes() *fyne.Container {

	lanesContainer := container.New(layout.NewGridLayout(6))

	for i := 1; i < 7; i++ {
		l := NewLane(i, a.Sw)
		lanesContainer.Add(l.Container)
		a.lanes = append(a.lanes, l)
	}

	return lanesContainer
}

func (a *App) startBtnPressed() {
	a.Sw.Start()
	go a.updateTimerText()
}

func (a *App) updateTimerText() {
	for time := range a.Sw.C {
		a.mainTimerText.Text = time
		a.mainTimerText.Refresh()
	}
}

func (a *App) stopBtnPressed() {
	a.Sw.Stop()
}

type Lane struct {
	Number    int
	sw        *stop_watch.StopWatch
	text      *canvas.Text
	Container *fyne.Container
}

func NewLane(laneNumber int, sw *stop_watch.StopWatch) *Lane {
	l := Lane{
		sw:     sw,
		Number: laneNumber,
	}

	l.createLane()

	return &l
}

func (l *Lane) createLane() {
	laneTitle := canvas.NewText(fmt.Sprintf("Bana %d", l.Number), white)
	l.text = canvas.NewText("00.00.00", white)
	laneTitle.TextSize = 20
	l.text.TextSize = 20
	vbox := container.New(layout.NewVBoxLayout(), laneTitle, l.text)
	l.Container = container.New(layout.NewHBoxLayout(), layout.NewSpacer(), vbox, layout.NewSpacer())
}
