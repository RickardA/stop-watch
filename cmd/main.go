package main

import (
	"github.com/RickardA/stop-watch/internal/app/gui"
)

func main() {
	// port := "8081"

	// ctx, cancelFunc := context.WithCancel(context.Background())
	// defer cancelFunc()

	// Setup Client
	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	gui.NewApp("Stop Watch")
}

// var white = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

// type Lane struct {
// 	number int
// 	text   *canvas.Text
// 	time   int
// 	sw     *StopWatch
// }

// func (l *Lane) StopTime() {
// 	fmt.Println("Stop timsss")
// 	l.time = l.sw.GetCurrentCount()
// 	l.text.Text = msToTime(l.time)
// 	l.text.Refresh()
// }

// func main() {
// 	lanes := map[int]*Lane{}
// 	sw := StopWatch{}

// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Hello")

// 	content := container.New(layout.NewGridLayout(6))

// 	for i := 1; i <= 6; i++ {
// 		container, lane := createLaneContainer(i)
// 		content.Add(container)

// 		lane.sw = &sw
// 		lanes[lane.number] = lane
// 	}

// 	myWindow.Resize(fyne.NewSize(800, 300))

// 	timerText, topBarContainer := createTopBarContainer()

// 	bottomBarContainer := createBottomBarContainer(func() {
// 		// Start
// 		sw.timer = time.NewTicker(time.Millisecond)

// 		go func(timerText *canvas.Text, sw *StopWatch, lanes map[int]*Lane) {
// 			sw.counter = 0
// 			for range sw.timer.C {
// 				sw.counter += 1
// 				timerText.Text = msToTime(sw.counter)
// 				timerText.Refresh()

// 				if sw.counter == 10000 {
// 					lanes[1].StopTime()
// 				}
// 			}
// 		}(timerText, &sw, lanes)
// 	}, func() {
// 		sw.timer.Stop()
// 	})

// 	mainContainer := container.New(layout.NewBorderLayout(topBarContainer, bottomBarContainer, nil, nil),
// 		topBarContainer, bottomBarContainer, content)

// 	myWindow.SetContent(mainContainer)
// 	myWindow.Show()
// 	myApp.Run()
// }

// func createTopBarContainer() (*canvas.Text, *fyne.Container) {
// 	timerText := canvas.NewText("00:00:00", color.White)
// 	timerText.TextSize = 50
// 	return timerText, container.New(layout.NewHBoxLayout(), layout.NewSpacer(), timerText, layout.NewSpacer())
// }

// func createBottomBarContainer(startCallback func(), stopCallback func()) *fyne.Container {
// 	startBtn := widget.NewButton("STARTA", startCallback)
// 	stopBtn := widget.NewButton("STANNA", stopCallback)
// 	btnContainer := container.New(layout.NewHBoxLayout(), startBtn, stopBtn)
// 	bottomBarContainer := container.New(layout.NewCenterLayout(), layout.NewSpacer(), btnContainer, layout.NewSpacer())

// 	return bottomBarContainer
// }

// func createLaneContainer(laneNumber int) (*fyne.Container, *Lane) {
// 	laneTitle := canvas.NewText(fmt.Sprintf("Bana %d", laneNumber), white)
// 	timeText := canvas.NewText("00.00.00", white)
// 	laneTitle.TextSize = 20
// 	timeText.TextSize = 20
// 	vbox := container.New(layout.NewVBoxLayout(), laneTitle, timeText)
// 	hbox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), vbox, layout.NewSpacer())

// 	l := Lane{
// 		number: laneNumber,
// 		text:   timeText,
// 		time:   0,
// 	}

// 	return hbox, &l
// }
