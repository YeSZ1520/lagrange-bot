package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type LifeCycleHandle func(model.LifeCycle, *Context)

func (receiver *Receiver) RegisterLifeCycleHandle(handle LifeCycleHandle) {
	if receiver.lifeCycleHandle == nil {
		receiver.lifeCycleHandle = make([]LifeCycleHandle, 0)
	}
	receiver.lifeCycleHandle = append(receiver.lifeCycleHandle, handle)
}

func (receiver *Receiver) LifeCycleRouter(msg model.LifeCycle, c *Context) {
	for _, handle := range receiver.lifeCycleHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
