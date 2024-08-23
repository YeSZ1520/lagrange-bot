package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

// group_recall_notice_handle

type GroupRecallNoticeHandle func(model.GroupRecallNotice, *Context)

func (receiver *Receiver) RegisterGroupRecallNoticeHandle(handle GroupRecallNoticeHandle) {
	if receiver.groupRecallNoticeHandle == nil {
		receiver.groupRecallNoticeHandle = make([]GroupRecallNoticeHandle, 0)
	}
	receiver.groupRecallNoticeHandle = append(receiver.groupRecallNoticeHandle, handle)
}

func (receiver *Receiver) GroupRecallNoticeRouter(msg model.GroupRecallNotice, c *Context) {
	for _, handle := range receiver.groupRecallNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
