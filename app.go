package kkapp

import (
	"github.com/kkserver/kk-lib/app"
)

const KKAppNameKey = "appname"

func New(parent app.IApp) *app.App {

	var v = app.NewApp(parent)

	v.Service(&KKService{})(&KKConnectTask{}, &KKDisconnectTask{}, &KKSendMessageTask{}, &KKReciveMessageTask{})

	v.Service(&KKRemoteService{})(&KKRemoteConnectTask{}, &KKRemoteTask{})

	return v
}
