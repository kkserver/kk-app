package kkapp

import (
	"github.com/kkserver/kk-lib/app"
	"time"
)

type KKRemoteTask struct {
	app.Task
	Name    string
	APITask app.IAPITask
	Timeout time.Duration
}
