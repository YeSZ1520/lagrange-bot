package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type GroupIncreaseNoticeHandle func(model.GroupIncreaseNotice, *Context)

func (receiver *Receiver) RegisterGroupIncreaseNoticeHandle(handle GroupIncreaseNoticeHandle) {
	if receiver.groupIncreaseNoticeHandle == nil {
		receiver.groupIncreaseNoticeHandle = make([]GroupIncreaseNoticeHandle, 0)
	}
	receiver.groupIncreaseNoticeHandle = append(receiver.groupIncreaseNoticeHandle, handle)
}

func (receiver *Receiver) GroupIncreaseNoticeRouter(msg model.GroupIncreaseNotice, c *Context) {
	for _, handle := range receiver.groupIncreaseNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
