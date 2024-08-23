package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type GroupRequestHandle func(model.GroupRequest, *Context)

func (receiver *Receiver) RegisterGroupRequestHandle(handle GroupRequestHandle) {
	if receiver.groupRequestHandle == nil {
		receiver.groupRequestHandle = make([]GroupRequestHandle, 0)
	}
	receiver.groupRequestHandle = append(receiver.groupRequestHandle, handle)
}

func (receiver *Receiver) GroupRequestRouter(msg model.GroupRequest, c *Context) {
	for _, handle := range receiver.groupRequestHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
