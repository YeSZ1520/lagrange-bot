# lagrange-bot

基于 [Lagrange.Core](https://github.com/LagrangeDev/Lagrange.Core) 的 Go 语言实现机器人框架，实现了 Lagrange.Core 的拓展
API ，使用Http Post 通讯

# API

## 使用

1. 创建 `api.Api` 并传入 Lagrange.Core 地址
2. 调用 API

```go
api := apis.Api{Hosts: "http://lagrange-core:8081"}
api.SendPrivateMsg(userID, message, false)
```

## OneBot V11 API

| API                      | Method               |
|--------------------------|----------------------|
| /send_private_msg        | SendPrivateMsg       |
| /send_group_msg          | SendGroupMessage     |
| /send_msg                | SendMsg              |
| /delete_msg              | DeleteMsg            |
| /get_msg                 | GetMsg               |
| /get_forward_msg         | GetForwardMsg        |
| /send_like               | SendLike             |
| /set_group_kick          | SetGroupKick         |
| /set_group_ban           | SetGroupBan          |
| /set_group_whole_ban     | SetGroupWholeBan     |
| /set_group_admin         | SetGroupAdmin        |
| /set_group_card          | SetGroupCard         |
| /set_group_name          | SetGroupName         |
| /set_group_leave         | SetGroupLeave        |
| /set_group_special_title | SetGroupSpecialTitle |
| /set_friend_add_request  | SetFriendAddRequest  |
| /set_group_add_request   | SetGroupAddRequest   |
| /get_login_info          | GetLoginInfo         |
| /get_stranger_info       | GetStrangerInfo      |
| /get_friend_list         | GetFriendList        |
| /get_group_info          | GetGroupInfo         |
| /get_group_list          | GetGroupList         |
| /get_group_member_info   | GetGroupMemberInfo   |
| /get_group_member_list   | GetGroupMemberList   |
| /get_group_honor_info    | GetGroupHonorInfo    |
| /get_cookies             | GetCookies           |
| /get_csrf_token          | GetCsrfToken         |
| /get_credentials         | GetCredentials       |
| /can_send_image          | CanSendImage         |
| /can_send_record         | CanSendRecord        |
| /get_status              | GetStatus            |
| /get_version_info        | GetVersionInfo       |
| /set_restart             | SetRestart           |

## 拓展 API

| API                        | Method                |
|----------------------------|-----------------------|
| /fetch_custom_face         | FetchCustomFace       |
| /get_friend_msg_history    | GetFriendMsgHistory   |
| /get_group_msg_history     | GetGroupMsgHistory    |
| /send_forward_msg          | SendForwardMsg        |
| /send_group_forward_msg    | SendGroupForwardMsg   |
| /send_private_forward_msg  | SendPrivateForwardMsg |
| /upload_group_file         | UploadGroupFile       |
| /upload_private_file       | UploadPrivateFile     |
| /get_group_root_files      | GetGroupRootFiles     |
| /get_group_files_by_folder | GetGroupFilesByFolder |
| /get_group_file_url        | GetGroupFilesByFolder |
| /friend_poke               | FriendPoke            |
| /group_poke                | GroupPoke             |

# 事件

## 使用

1. 创建 `receivers.Receiver`
2. 注册事件处理函数
3. 调用 `Run` 启动服务

```go
r := receivers.Receiver{}
r.RegisterPrivateMessageHandle(func (message model.PrivateMessage, context *receivers.Context) {
    time.Sleep(5 * time.Second)
    if message.Sender.UserID != userID {
        context.Abort()
    }
})

r.RegisterPrivateMessageHandle(func (message model.PrivateMessage, context *receivers.Context) {
    _, err := api.SendPrivateMsg(message.Sender.UserID, model.ToCQCodes(message.Message), false)
    if err != nil {
        slog.Error("Send Private Msg Error" + err.Error())
    }
})
r.Run(":8745")
```

+ 事件注册会按注册先后顺序执行
+ 若调用 `Context.Abort()` 则停止执行后续事件处理函数
+ 可通过 `Context` 在事件处理函数之间传递数据

## 支持

| 类型 | 事件      | 事件注册函数                            |
|----|---------|-----------------------------------|
| 消息 | 私聊消息    | RegisterPrivateMessageHandle      |
| 消息 | 群消息     | RegisterGroupMessageHandle        | 
| 通知 | 群文件上传   | RegisterGroupUploadNoticeHandle   | 
| 通知 | 群管理员变动  | RegisterGroupAdminNoticeHandle    | 
| 通知 | 群成员减少   | RegisterGroupDecreaseNoticeHandle | 
| 通知 | 群成员增加   | RegisterGroupIncreaseNoticeHandle | 
| 通知 | 群禁言     | RegisterGroupBanNoticeHandle      |
| 通知 | 好友添加    | RegisterFriendAddNoticeHandle     | 
| 通知 | 群消息撤回   | RegisterGroupRecallNoticeHandle   | 
| 通知 | 好友消息撤回  | RegisterFriendRecallNoticeHandle  |
| 请求 | 加好友请求   | RegisterFriendRequestHandle       | 
| 请求 | 加群请求／邀请 | RegisterGroupRequestHandle        | 
| 元  | 生命周期    | RegisterLifeCycleHandle           | 
| 元  | 心跳事件    | RegisterHeartBeatHandle           | 

# 参考

## 参考 docker-compose.yml

```yaml
services:
  lagrange-core:
    image: ghcr.io/lagrangedev/lagrange.onebot:edge
    restart: always
    container_name: lagrange-core
    volumes:
      - /opt/lagrange:/app/data
      - /opt/lagrange/upload:/app/upload
    environment:
      - TZ=Asia/Shanghai

  github.com/YeSZ1520/lagrange-bot:
    build:
      context: .
    image: github.com/YeSZ1520/lagrange-bot:1.6
    restart: always
    container_name: github.com/YeSZ1520/lagrange-bot
    volumes:
      - /opt/lagrange/upload:/app/upload
```

## 参考 appsettings.json

```json
[
  {
    "Type": "HttpPost",
    "Host": "lagrange-bot",
    "Port": 8745,
    "Suffix": "/",
    "HeartBeatInterval": 20000,
    "HeartBeatEnable": true,
    "AccessToken": ""
  },
  {
    "Type": "Http",
    "Host": "*",
    "Port": 8081,
    "AccessToken": ""
  }
]
```