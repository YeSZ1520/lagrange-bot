package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type HeartBeatHandle func(model.HeartBeat, *Context)

func (receiver *Receiver) RegisterHeartBeatHandle(handle HeartBeatHandle) {
	if receiver.heartBeatHandle == nil {
		receiver.heartBeatHandle = make([]HeartBeatHandle, 0)
	}
	receiver.heartBeatHandle = append(receiver.heartBeatHandle, handle)
}

func (receiver *Receiver) HeartBeatRouter(msg model.HeartBeat, c *Context) {
	for _, handle := range receiver.heartBeatHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
