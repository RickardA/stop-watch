package app

import (
	"github.com/RickardA/stop-watch/internal/app/gui"
	"github.com/RickardA/stop-watch/internal/pkg/stop_watch"
)

type Client struct {
	Sw  *stop_watch.StopWatch
	App *gui.App
}

func NewClient() Client {
	return Client{
		Sw:  stop_watch.NewStopWatch(),
		App: gui.NewApp("Stop Watch"),
	}
}
