package receivers

import "github.com/YeSZ1520/lagrange-bot/model"

type FriendRequestHandle func(model.FriendRequest, *Context)

func (receiver *Receiver) RegisterFriendRequestHandle(handle FriendRequestHandle) {
	if receiver.friendRequestHandle == nil {
		receiver.friendRequestHandle = make([]FriendRequestHandle, 0)
	}
	receiver.friendRequestHandle = append(receiver.friendRequestHandle, handle)
}

func (receiver *Receiver) FriendRequestRouter(msg model.FriendRequest, c *Context) {
	for _, handle := range receiver.friendRequestHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
