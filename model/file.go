package model

type Files struct {
	Files   []File   `json:"files"`
	Folders []Folder `json:"folders"`
}

type File struct {
	GroupID       int64  `json:"group_id"`
	FileID        string `json:"file_id"`
	FileName      string `json:"file_name"`
	BusID         int64  `json:"busid"`
	FileSize      int64  `json:"file_size"`
	UploadTime    int64  `json:"upload_time"`
	DeadTime      int64  `json:"dead_time"`
	ModifyTime    int64  `json:"modify_time"`
	DownloadTimes int64  `json:"download_times"`
	Uploader      int64  `json:"uploader"`
	UploaderName  string `json:"uploader_name"`
}

type Folder struct {
	GroupID        int64  `json:"group_id"`
	FolderID       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	CreateTime     int64  `json:"create_time"`
	Creator        int64  `json:"creator"`
	CreatorName    string `json:"creator_name"`
	TotalFileCount int64  `json:"total_file_count"`
}

type FileUrl struct {
	Url string `json:"url"`
}
