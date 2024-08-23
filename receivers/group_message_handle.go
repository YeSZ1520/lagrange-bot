package receivers

import (
	"github.com/YeSZ1520/lagrange-bot/model"
)

// GroupMessageHandle 群聊消息处理
type GroupMessageHandle func(model.GroupMessage, *Context)

// RegisterGroupMessageHandle 注册群聊消息处理
func (receiver *Receiver) RegisterGroupMessageHandle(handle GroupMessageHandle) {
	if receiver.groupMessageHandle == nil {
		receiver.groupMessageHandle = make([]GroupMessageHandle, 0)
	}
	receiver.groupMessageHandle = append(receiver.groupMessageHandle, handle)
}

// GroupMessageRouter 群聊消息路由
func (receiver *Receiver) GroupMessageRouter(msg model.GroupMessage, c *Context) {
	for _, handle := range receiver.groupMessageHandle {
		if c.IsAbort() {
			return
		}
		handle(msg, c)
	}
}
