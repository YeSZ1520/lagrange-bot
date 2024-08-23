package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type GroupAdminNoticeHandle func(model.GroupAdminNotice, *Context)

func (receiver *Receiver) RegisterGroupAdminNoticeHandle(handle GroupAdminNoticeHandle) {
	if receiver.groupAdminNoticeHandle == nil {
		receiver.groupAdminNoticeHandle = make([]GroupAdminNoticeHandle, 0)
	}
	receiver.groupAdminNoticeHandle = append(receiver.groupAdminNoticeHandle, handle)
}

func (receiver *Receiver) GroupAdminNoticeRouter(msg model.GroupAdminNotice, c *Context) {
	for _, handle := range receiver.groupAdminNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
