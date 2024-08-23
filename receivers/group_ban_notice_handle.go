package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type GroupBanNoticeHandle func(model.GroupBanNotice, *Context)

func (receiver *Receiver) RegisterGroupBanNoticeHandle(handle GroupBanNoticeHandle) {
	if receiver.groupBanNoticeHandle == nil {
		receiver.groupBanNoticeHandle = make([]GroupBanNoticeHandle, 0)
	}
	receiver.groupBanNoticeHandle = append(receiver.groupBanNoticeHandle, handle)
}

func (receiver *Receiver) GroupBanNoticeRouter(msg model.GroupBanNotice, c *Context) {
	for _, handle := range receiver.groupBanNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
