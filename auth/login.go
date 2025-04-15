package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// GetLoginURL 通过用户名密码获取token, 并生成登录URL
func GetLoginURL(loginName, password, loginTypeId, univCode string) (string, error) {
	token, err := getToken(loginName, password)
	if err != nil {
		return "", err
	}
	loginURL := fmt.Sprintf(
		"https://vub-3f3ab907.gaojidata.com/login?loginTypeId=%s&univCode=%s&loginName=%s&token=%s",
		loginTypeId, univCode, loginName, token,
	)
	return loginURL, nil
}

// getToken 执行loginCG请求, 获取token
func getToken(loginName, password string) (string, error) {
	// 字节缓冲区, 用来存放构造的请求体
	var buf bytes.Buffer
	// multipart写入器, 可以把key-value键值对写入buffer并自动生成multipart边界信息
	writer := multipart.NewWriter(&buf)
	writer.WriteField("LoginName", loginName)
	writer.WriteField("Password", password)
	writer.Close()

	req, err := http.NewRequest("POST", "https://cg-3f3ab907.gaojidata.com/api/v1/sso/login/loginByPwd", &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json, text/plain, */*; charset=utf-8")
	req.Header.Set("Referer", "https://cg-3f3ab907.gaojidata.com/login")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//fmt.Println("原始响应体：", string(body))

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	data, ok := result["data"].(string)
	if !ok {
		return "", fmt.Errorf("getToken fail")
	}
	return data, nil
}
