package kkapp

import (
	"encoding/json"
	"errors"
	"github.com/kkserver/kk-lib/app"
	"github.com/kkserver/kk-lib/kk"
	"time"
)

type KKRemoteService struct {
	app.Service
	request func(message *kk.Message, timeout time.Duration) *kk.Message
	getName func() string
}

func (S *KKRemoteService) Handle(a app.IApp, task app.ITask) error {
	return S.ReflectHandle(a, task, S)
}

func (S *KKRemoteService) HandleKKRemoteTask(a app.IApp, task *KKRemoteTask) error {

	if S.request != nil {

		var v = kk.Message{}

		v.To = task.Name + task.APITask.API()
		v.Type = "text/json"
		v.Content, _ = json.Marshal(task.APITask)

		var r = S.request(&v, task.Timeout)

		if r == nil {
			return errors.New("KKRemoteService request fail")
		} else if r.Method == "REQUEST" && (r.Type == "text/json" || r.Type == "application/json") {
			return json.Unmarshal(r.Content, task.APITask.GetResult())
		} else {
			return errors.New("KKRemoteService request fail " + r.String())
		}

	} else {
		return errors.New("KKRemoteService not connected")
	}

	return nil
}

func (S *KKRemoteService) HandleKKRemoteConnectTask(a app.IApp, task *KKRemoteConnectTask) error {

	if S.request == nil {
		S.request, S.getName = kk.TCPClientRequestConnect(task.Name, task.Address, task.Options)
	}

	return nil
}
