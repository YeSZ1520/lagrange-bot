package apis

import (
	"fmt"
	"github.com/YeSZ1520/lagrange-bot/model"
)

func (api Api) SendPrivateMsg(uid int64, msg string, autoEscape bool) (*model.ApiResult[model.MessageId], error) {
	if msg == "" {
		return nil, fmt.Errorf("msg is empty")
	}
	body := map[string]interface{}{
		"user_id":     uid,
		"message":     msg,
		"auto_escape": autoEscape,
	}
	result, err := SendApiRequest[model.MessageId](api.Hosts+"/send_private_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) SendGroupMessage(gid int64, msg string, autoEscape bool) (*model.ApiResult[model.MessageId], error) {
	if msg == "" {
		return nil, fmt.Errorf("msg is empty")
	}
	body := map[string]interface{}{
		"group_id":    gid,
		"message":     msg,
		"auto_escape": autoEscape,
	}
	result, err := SendApiRequest[model.MessageId](api.Hosts+"/send_group_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) SendMsg(msgType string, uid int64, gid int64, msg string, autoEscape bool) (*model.ApiResult[model.MessageId], error) {
	if msg == "" {
		return nil, fmt.Errorf("msg is empty")
	}
	body := map[string]interface{}{
		"message_type": msgType,
		"user_id":      uid,
		"group_id":     gid,
		"message":      msg,
		"auto_escape":  autoEscape,
	}
	result, err := SendApiRequest[model.MessageId](api.Hosts+"/send_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) SendForwardMsg(msg []model.NodeSegmentData) (*model.ApiResult[string], error) {
	if len(msg) == 0 {
		return nil, fmt.Errorf("msg is empty")
	}
	var messages []model.NodeSegment
	for _, message := range msg {
		messages = append(messages, model.NodeSegment{
			Type: "node",
			Data: message,
		})
	}
	body := map[string]interface{}{
		"messages": messages,
	}
	result, err := SendApiRequest[string](api.Hosts+"/send_forward_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) SendPrivateForwardMsg(uid int64, msg []model.NodeSegmentData) (*model.ApiResult[model.ForwardResult], error) {
	if len(msg) == 0 {
		return nil, fmt.Errorf("msg is empty")
	}
	var messages []model.NodeSegment
	for _, message := range msg {
		messages = append(messages, model.NodeSegment{
			Type: "node",
			Data: message,
		})
	}
	body := map[string]interface{}{
		"user_id":  uid,
		"messages": messages,
	}
	result, err := SendApiRequest[model.ForwardResult](api.Hosts+"/send_private_forward_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) SendGroupForwardMsg(gid int64, msg []model.NodeSegmentData) (*model.ApiResult[model.ForwardResult], error) {
	if len(msg) == 0 {
		return nil, fmt.Errorf("msg is empty")
	}
	var messages []model.NodeSegment
	for _, message := range msg {
		messages = append(messages, model.NodeSegment{
			Type: "node",
			Data: message,
		})
	}
	body := map[string]interface{}{
		"group_id": gid,
		"messages": messages,
	}
	result, err := SendApiRequest[model.ForwardResult](api.Hosts+"/send_group_forward_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) DeleteMsg(msgID int64) error {
	body := map[string]int64{
		"message_id": msgID,
	}
	_, err := SendApiRequest[map[string]int64](api.Hosts+"/delete_msg", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) GetMsg(msgID int64) (*model.ApiResult[model.Message], error) {
	body := map[string]int64{
		"message_id": msgID,
	}
	result, err := SendApiRequest[model.Message](api.Hosts+"/get_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetForwardMsg(msgID string) (*model.ApiResult[[]model.Segment], error) {
	body := map[string]string{
		"id": msgID,
	}
	result, err := SendApiRequest[[]model.Segment](api.Hosts+"/get_forward_msg", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) FetchCustomFace() (*model.ApiResult[[]string], error) {
	result, err := SendApiRequest[[]string](api.Hosts+"/fetch_custom_face", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetFriendMsgHistory(uid int64, messageID int64, count int) (*model.ApiResult[[]model.Segment], error) {
	body := map[string]interface{}{
		"user_id":    uid,
		"message_id": messageID,
		"count":      count,
	}
	result, err := SendApiRequest[[]model.Segment](api.Hosts+"/get_friend_msg_history", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupMsgHistory(gid int64, messageID int64, count int) (*model.ApiResult[[]model.Segment], error) {
	body := map[string]interface{}{
		"group_id":   gid,
		"message_id": messageID,
		"count":      count,
	}
	result, err := SendApiRequest[[]model.Segment](api.Hosts+"/get_group_msg_history", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) FriendPoke(uid int64) error {
	body := map[string]int64{
		"user_id": uid,
	}
	_, err := SendApiRequest[map[string]int64](api.Hosts+"/friend_poke", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) GroupPoke(gid int64, uid int64) error {
	body := map[string]int64{
		"group_id": gid,
		"user_id":  uid,
	}
	_, err := SendApiRequest[map[string]int64](api.Hosts+"/group_poke", body)
	if err != nil {
		return err
	}
	return nil
}
