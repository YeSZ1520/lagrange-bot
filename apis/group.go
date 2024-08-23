package apis

import "github.com/YeSZ1520/lagrange-bot/model"

func (api Api) SetGroupKick(gid int64, uid int64, rejectAddRequest bool) error {
	body := map[string]interface{}{
		"group_id":           gid,
		"user_id":            uid,
		"reject_add_request": rejectAddRequest,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_kick", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupBan(gid int64, uid int64, duration int64) error {
	body := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"duration": duration,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_ban", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupWholeBan(gid int64, enable bool) error {
	body := map[string]interface{}{
		"group_id": gid,
		"enable":   enable,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_whole_ban", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupAdmin(gid int64, uid int64, enable bool) error {
	body := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"enable":   enable,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_admin", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupCard(gid int64, uid int64, card string) error {
	body := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"card":     card,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_card", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupName(gid int64, name string) error {
	body := map[string]interface{}{
		"group_id":   gid,
		"group_name": name,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_name", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupLeave(gid int64, isDismiss bool) error {
	body := map[string]interface{}{
		"group_id":   gid,
		"is_dismiss": isDismiss,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_leave", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupSpecialTitle(gid int64, uid int64, title string, duration int64) error {
	body := map[string]interface{}{
		"group_id":      gid,
		"user_id":       uid,
		"special_title": title,
		"duration":      duration,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_special_title", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) SetGroupAddRequest(flag string, subType string, approve bool, reason string) error {
	body := map[string]interface{}{
		"flag":     flag,
		"sub_type": subType,
		"approve":  approve,
		"reason":   reason,
	}
	_, err := SendApiRequest[any](api.Hosts+"/set_group_add_request", body)
	if err != nil {
		return err
	}
	return nil
}

func (api Api) GetGroupList() (*model.ApiResult[[]model.GroupInfo], error) {
	result, err := SendApiRequest[[]model.GroupInfo](api.Hosts+"/get_group_list", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupMemberList(gid int64) (*model.ApiResult[[]model.GroupMemberInfo], error) {
	body := map[string]interface{}{
		"group_id": gid,
	}
	result, err := SendApiRequest[[]model.GroupMemberInfo](api.Hosts+"/get_group_member_list", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupMemberInfo(gid int64, uid int64, noCache bool) (*model.ApiResult[model.GroupMemberInfo], error) {
	body := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"no_cache": noCache,
	}
	result, err := SendApiRequest[model.GroupMemberInfo](api.Hosts+"/get_group_member_info", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupHonorInfo(gid int64, type_ string) (*model.ApiResult[model.GroupHonorInfo], error) {
	body := map[string]interface{}{
		"group_id": gid,
		"type":     type_,
	}
	result, err := SendApiRequest[model.GroupHonorInfo](api.Hosts+"/get_group_honor_info", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
