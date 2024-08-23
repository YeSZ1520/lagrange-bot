package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/YeSZ1520/lagrange-bot/model"
	"net/http"
)

type Api struct {
	Hosts string `json:"hosts"`
}

func ParseResult[T any](resp *http.Response) (*model.ApiResult[T], error) {
	var data model.ApiResult[T]
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if data.Status != "ok" {
		return nil, fmt.Errorf("API returned status %s", data.Status)
	}
	return &data, nil
}

func SendApiRequest[T any](url string, body interface{}) (*model.ApiResult[T], error) {
	var jsonBody []byte = nil
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}
	return ParseResult[T](resp)
}
func (api Api) SetRestart() error {
	_, err := SendApiRequest[any](api.Hosts+"/set_restart", nil)
	if err != nil {
		return err
	}
	return nil
}
