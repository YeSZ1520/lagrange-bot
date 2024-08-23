package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type GroupDecreaseNoticeHandle func(model.GroupDecreaseNotice, *Context)

func (receiver *Receiver) RegisterGroupDecreaseNoticeHandle(handle GroupDecreaseNoticeHandle) {
	if receiver.groupDecreaseNoticeHandle == nil {
		receiver.groupDecreaseNoticeHandle = make([]GroupDecreaseNoticeHandle, 0)
	}
	receiver.groupDecreaseNoticeHandle = append(receiver.groupDecreaseNoticeHandle, handle)
}

func (receiver *Receiver) GroupDecreaseNoticeRouter(msg model.GroupDecreaseNotice, c *Context) {
	for _, handle := range receiver.groupDecreaseNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
