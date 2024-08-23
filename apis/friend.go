package apis

import "github.com/YeSZ1520/lagrange-bot/model"

// SendLike 发送好友赞
func (api Api) SendLike(uid int64, times int64) error {
	body := map[string]interface{}{
		"user_id": uid,
		"times":   times,
	}
	_, err := SendApiRequest[map[string]int64](api.Hosts+"/send_like", body)
	if err != nil {
		return err
	}
	return nil
}

// SetFriendAddRequest 处理加好友请求
func (api Api) SetFriendAddRequest(flag string, approve bool, remark string) error {
	body := map[string]interface{}{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	}
	_, err := SendApiRequest[map[string]int64](api.Hosts+"/set_friend_add_request", body)
	if err != nil {
		return err
	}
	return nil
}

// GetFriendList 获取好友列表
func (api Api) GetFriendList() (*model.ApiResult[model.FriendInfo], error) {
	result, err := SendApiRequest[model.FriendInfo](api.Hosts+"/get_friend_list", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStrangerInfo 获取陌生人信息
func (api Api) GetStrangerInfo(uid int64, noCache bool) (*model.ApiResult[model.UserInfo], error) {
	body := map[string]interface{}{
		"user_id":  uid,
		"no_cache": noCache,
	}
	result, err := SendApiRequest[model.UserInfo](api.Hosts+"/get_stranger_info", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
