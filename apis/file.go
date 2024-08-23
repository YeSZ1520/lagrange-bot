package apis

import "github.com/YeSZ1520/lagrange-bot/model"

func (api Api) UploadPrivateFile(uid int64, file string, name string) (*model.ApiResult[any], error) {
	body := map[string]interface{}{
		"user_id": uid,
		"file":    file,
		"name":    name,
	}
	result, err := SendApiRequest[any](api.Hosts+"/upload_private_file", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) UploadGroupFile(gid int64, file string, name string, folder string) (*model.ApiResult[any], error) {
	body := map[string]interface{}{
		"group_id": gid,
		"file":     file,
		"name":     name,
		"folder":   folder,
	}
	result, err := SendApiRequest[any](api.Hosts+"/upload_group_file", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupRootFiles(gid int64) (*model.ApiResult[model.Files], error) {
	body := map[string]interface{}{
		"group_id": gid,
	}
	result, err := SendApiRequest[model.Files](api.Hosts+"/get_group_root_files", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupFilesByFolder(gid int64, folderID string) (*model.ApiResult[model.Files], error) {
	body := map[string]interface{}{
		"group_id":  gid,
		"folder_id": folderID,
	}
	result, err := SendApiRequest[model.Files](api.Hosts+"/get_group_files_by_folder", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (api Api) GetGroupFileUrl(gid int64, fileID string, busid int64) (*model.ApiResult[model.FileUrl], error) {
	body := map[string]interface{}{
		"group_id": gid,
		"file_id":  fileID,
		"busid":    busid,
	}
	result, err := SendApiRequest[model.FileUrl](api.Hosts+"/get_group_file_url", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
