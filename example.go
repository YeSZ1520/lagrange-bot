package main

import (
	"github.com/YeSZ1520/lagrange-bot/apis"
	"github.com/YeSZ1520/lagrange-bot/model"
	"github.com/YeSZ1520/lagrange-bot/receivers"
	"log/slog"
	"time"
)

func main() {
	var userID int64 = 12345678
	api := apis.Api{Hosts: "http://lagrange-core:8081"}

	r := receivers.Receiver{}
	r.RegisterPrivateMessageHandle(func(message model.PrivateMessage, context *receivers.Context) {
		time.Sleep(5 * time.Second)
		println(message.RawMessage)
		if message.Sender.UserID != userID {
			context.Abort()
		}
	})

	r.RegisterPrivateMessageHandle(func(message model.PrivateMessage, context *receivers.Context) {
		_, err := api.SendPrivateMsg(message.Sender.UserID, model.ToCQCodes(message.Message), false)
		if err != nil {
			slog.Error("Send Private Msg Error" + err.Error())
			return
		}
	})
	r.Run(":8745")

}
