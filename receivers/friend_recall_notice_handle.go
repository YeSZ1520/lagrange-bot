package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type FriendRecallNoticeHandle func(model.FriendRecallNotice, *Context)

func (receiver *Receiver) RegisterFriendRecallNoticeHandle(handle FriendRecallNoticeHandle) {
	if receiver.friendRecallNoticeHandle == nil {
		receiver.friendRecallNoticeHandle = make([]FriendRecallNoticeHandle, 0)
	}
	receiver.friendRecallNoticeHandle = append(receiver.friendRecallNoticeHandle, handle)
}

func (receiver *Receiver) FriendRecallNoticeRouter(msg model.FriendRecallNotice, c *Context) {
	for _, handle := range receiver.friendRecallNoticeHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
