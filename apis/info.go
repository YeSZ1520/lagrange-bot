package apis

import "github.com/YeSZ1520/lagrange-bot/model"

func (api Api) GetLoginInfo() (*model.ApiResult[model.LoginInfo], error) {
	result, err := SendApiRequest[model.LoginInfo](api.Hosts+"/get_login_info", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupInfo(gid int64, noCache bool) (*model.ApiResult[model.GroupInfo], error) {
	body := map[string]interface{}{
		"group_id": gid,
		"no_cache": noCache,
	}
	result, err := SendApiRequest[model.GroupInfo](api.Hosts+"/get_group_info", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) CanSendImage() (*model.ApiResult[model.CanSend], error) {
	result, err := SendApiRequest[model.CanSend](api.Hosts+"/can_send_image", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) CanSendRecord() (*model.ApiResult[model.CanSend], error) {
	result, err := SendApiRequest[model.CanSend](api.Hosts+"/can_send_record", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetStatus() (*model.ApiResult[model.LagrangeStatus], error) {
	result, err := SendApiRequest[model.LagrangeStatus](api.Hosts+"/get_status", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetVersionInfo() (*model.ApiResult[model.LagrangeVersionInfo], error) {
	result, err := SendApiRequest[model.LagrangeVersionInfo](api.Hosts+"/get_version_info", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
