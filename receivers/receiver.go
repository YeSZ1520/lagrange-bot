package receivers

import (
	"encoding/json"
	"fmt"
	"github.com/YeSZ1520/lagrange-bot/model"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
)

type Receiver struct {
	groupMessageHandle        []GroupMessageHandle
	privateMessageHandle      []PrivateMessageHandle
	groupUploadNoticeHandle   []GroupUploadNoticeHandle
	groupAdminNoticeHandle    []GroupAdminNoticeHandle
	groupDecreaseNoticeHandle []GroupDecreaseNoticeHandle
	groupIncreaseNoticeHandle []GroupIncreaseNoticeHandle
	groupBanNoticeHandle      []GroupBanNoticeHandle
	friendAddNoticeHandle     []FriendAddNoticeHandle
	groupRecallNoticeHandle   []GroupRecallNoticeHandle
	friendRecallNoticeHandle  []FriendRecallNoticeHandle
	friendRequestHandle       []FriendRequestHandle
	groupRequestHandle        []GroupRequestHandle
	heartBeatHandle           []HeartBeatHandle
	lifeCycleHandle           []LifeCycleHandle
}

type Context struct {
	Data map[string]any
	model.Aborts
}

func (receiver *Receiver) Run(addr ...string) {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("body err:", err)
			return
		}
		receiver.Router(body)
		c.String(204, "ok")
	})
	err := r.Run(addr...)
	if err != nil {
		return
	}
}

// Router 事件路由
func (receiver *Receiver) Router(body []byte) {
	// 推送事件
	var pushEvent model.PushEvent
	err := json.Unmarshal(body, &pushEvent)
	if err != nil {
		slog.Error("Push Event Json Parse Err:", err)
		return
	}
	context := &Context{
		Data:   map[string]any{},
		Aborts: model.Aborts{},
	}
	// event router
	switch pushEvent.PostType {
	// 消息事件
	case "message":
		if pushEvent.MessageType == "private" {
			// Private Message 私聊消息
			var privateMessage model.PrivateMessage
			err = json.Unmarshal(body, &privateMessage)
			receiver.PrivateMessageRouter(privateMessage, context)
		} else if pushEvent.MessageType == "group" {
			// Group Message 群聊消息
			var groupMessage model.GroupMessage
			err = json.Unmarshal(body, &groupMessage)
			receiver.GroupMessageRouter(groupMessage, context)
		}

	case "notice":
		slog.Debug("Notice Event", pushEvent.NoticeType)
		if pushEvent.NoticeType == "group_upload" {
			var groupUploadNotice model.GroupUploadNotice
			err = json.Unmarshal(body, &groupUploadNotice)
			receiver.GroupUploadNoticeRouter(groupUploadNotice, context)
		} else if pushEvent.NoticeType == "group_admin" {
			var groupAdminNotice model.GroupAdminNotice
			err = json.Unmarshal(body, &groupAdminNotice)
			receiver.GroupAdminNoticeRouter(groupAdminNotice, context)
		} else if pushEvent.NoticeType == "group_decrease" {
			var groupDecreaseNotice model.GroupDecreaseNotice
			err = json.Unmarshal(body, &groupDecreaseNotice)
			receiver.GroupDecreaseNoticeRouter(groupDecreaseNotice, context)
		} else if pushEvent.NoticeType == "group_increase" {
			var groupIncreaseNotice model.GroupIncreaseNotice
			err = json.Unmarshal(body, &groupIncreaseNotice)
			receiver.GroupIncreaseNoticeRouter(groupIncreaseNotice, context)
		} else if pushEvent.NoticeType == "group_ban" {
			var groupBanNotice model.GroupBanNotice
			err = json.Unmarshal(body, &groupBanNotice)
			receiver.GroupBanNoticeRouter(groupBanNotice, context)
		} else if pushEvent.NoticeType == "friend_add" {
			var friendAddNotice model.FriendAddNotice
			err = json.Unmarshal(body, &friendAddNotice)
			receiver.FriendAddNoticeRouter(friendAddNotice, context)
		} else if pushEvent.NoticeType == "group_recall" {
			var groupRecallNotice model.GroupRecallNotice
			err = json.Unmarshal(body, &groupRecallNotice)
			receiver.GroupRecallNoticeRouter(groupRecallNotice, context)
		} else if pushEvent.NoticeType == "friend_recall" {
			var friendRecallNotice model.FriendRecallNotice
			err = json.Unmarshal(body, &friendRecallNotice)
			receiver.FriendRecallNoticeRouter(friendRecallNotice, context)
		}

	case "request":
		if pushEvent.RequestType == "friend" {
			var friendRequest model.FriendRequest
			err = json.Unmarshal(body, &friendRequest)
			receiver.FriendRequestRouter(friendRequest, context)
		} else if pushEvent.RequestType == "group" {
			var groupRequest model.GroupRequest
			err = json.Unmarshal(body, &groupRequest)
			receiver.GroupRequestRouter(groupRequest, context)
		}
	case "meta_event":
		if pushEvent.MetaEventType == "heartbeat" {
			var heartBeat model.HeartBeat
			err = json.Unmarshal(body, &heartBeat)
			receiver.HeartBeatRouter(heartBeat, context)
		} else if pushEvent.MetaEventType == "lifecycle" {
			var lifeCycle model.LifeCycle
			err = json.Unmarshal(body, &lifeCycle)
			receiver.LifeCycleRouter(lifeCycle, context)
		}
	}
}
