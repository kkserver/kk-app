package kkapp

import (
	"github.com/kkserver/kk-lib/app"
)

type KKRemoteConnectTask struct {
	app.Task
	Name    string
	Address string
	Options map[string]interface{}
}
