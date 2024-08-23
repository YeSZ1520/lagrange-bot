package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type GroupUploadNoticeHandle func(model.GroupUploadNotice, *Context)

func (receiver *Receiver) RegisterGroupUploadNoticeHandle(handle GroupUploadNoticeHandle) {
	if receiver.groupUploadNoticeHandle == nil {
		receiver.groupUploadNoticeHandle = make([]GroupUploadNoticeHandle, 0)
	}
	receiver.groupUploadNoticeHandle = append(receiver.groupUploadNoticeHandle, handle)
}

func (receiver *Receiver) GroupUploadNoticeRouter(msg model.GroupUploadNotice, c *Context) {
	for _, handle := range receiver.groupUploadNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
