package httputils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func CurlWithParam(link string, method string, param map[string]string) (response string, err error) {
	client := http.Client{Timeout: time.Second * 10}
	m := strings.ToLower(method)
	var resp *http.Response

	if m == "get" {
		u, _ := url.Parse(link)
		q := u.Query()
		for k, val := range param {
			q.Set(k, val)
		}
		u.RawQuery = q.Encode()
		link = u.String()

		resp, err = client.Get(link)
	} else if m == "post" {
		str := ""
		for k, val := range param {
			if str == "" {
				str = k + "=" + val
			} else {
				str = str + "&" + k + "=" + val
			}
		}

		r := strings.NewReader(str)
		client := http.Client{Timeout: time.Second * 10}
		// resp, err = client.Post(link, "application/x-www-form-urlencoded", r)
		resp, err = client.Post(link, "application/json", r)
	} else {
		return response, errors.New("only support get and post method")
	}

	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求返回的内容
func Post(url string, data interface{}, contentType string) (content string) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	statusCode := resp.StatusCode
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("响应状态码为： %v\n", statusCode)
	content = string(result)
	return content
}
