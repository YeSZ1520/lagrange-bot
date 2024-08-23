package receivers

import (
	"github.com/YeSZ1520/lagrange-bot/model"
)

type PrivateMessageHandle func(model.PrivateMessage, *Context)

func (receiver *Receiver) RegisterPrivateMessageHandle(handle PrivateMessageHandle) {
	if receiver.privateMessageHandle == nil {
		receiver.privateMessageHandle = make([]PrivateMessageHandle, 0)
	}
	receiver.privateMessageHandle = append(receiver.privateMessageHandle, handle)
}

func (receiver *Receiver) PrivateMessageRouter(msg model.PrivateMessage, c *Context) {
	for _, handle := range receiver.privateMessageHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
