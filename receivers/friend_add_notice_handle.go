package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type FriendAddNoticeHandle func(model.FriendAddNotice, *Context)

func (receiver *Receiver) RegisterFriendAddNoticeHandle(handle FriendAddNoticeHandle) {
	if receiver.friendAddNoticeHandle == nil {
		receiver.friendAddNoticeHandle = make([]FriendAddNoticeHandle, 0)
	}
	receiver.friendAddNoticeHandle = append(receiver.friendAddNoticeHandle, handle)
}

func (receiver *Receiver) FriendAddNoticeRouter(msg model.FriendAddNotice, c *Context) {
	for _, handle := range receiver.friendAddNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
