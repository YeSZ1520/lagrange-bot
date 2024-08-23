package apis

import "github.com/YeSZ1520/lagrange-bot/model"

// GetCookies 获取Cookies
func (api Api) GetCookies(domain string) (*model.ApiResult[model.Cookies], error) {
	body := map[string]interface{}{
		"domain": domain,
	}
	result, err := SendApiRequest[model.Cookies](api.Hosts+"/get_cookies", body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCsrfToken 获取CsrfToken
func (api Api) GetCsrfToken() (*model.ApiResult[model.CsrfToken], error) {
	result, err := SendApiRequest[model.CsrfToken](api.Hosts+"/get_csrf_token", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCredentials 获取Credentials
func (api Api) GetCredentials() (*model.ApiResult[model.Credentials], error) {
	result, err := SendApiRequest[model.Credentials](api.Hosts+"/get_credentials", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
