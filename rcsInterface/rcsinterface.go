package rcsinterface

import (
	"fmt"
	"net/http"
	"strings"

	httputils "github.com/JasonFxvic/huike/httpUtils"
)

var BaseUrl string = "http://10.22.224.141:6546/"

//查询单台车辆信息
func GetAgvMessage(id string) (response string, err error) {
	targetUrl := BaseUrl + "api/v2/agvs/" + id
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}

//获取所有机器人信息
func GetAllAgvMessage() (response string, err error) {
	targetUrl := BaseUrl + "api/v2/agvs"
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}

//获取所有车辆正在执行的指令
func GetAgvOrders(id string) (response string, err error) {
	targetUrl := BaseUrl + "api/v2/agvs/" + id + "/orders"
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}

// 修改车辆参数
func UpdateAgvAttr(id string, data string) {
	targetUrl := BaseUrl + "api/v2/agvs/" + id
	payload := strings.NewReader(data)
	req, _ := http.NewRequest("PATCH", targetUrl, payload)
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(response)
		return
	}
	fmt.Print(response)
}

// 创建指令
func CreateOrder(data interface{}) {
	targetUrl := BaseUrl + "api/v2/orders"
	content := httputils.Post(targetUrl, data, "application/json")
	fmt.Print(content)
}

//按照id查询指令
func GetOrderById(id string) (response string, err error) {
	targetUrl := BaseUrl + "api/v2/orders/" + id
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}

//按照id查询多个指令
func GetOrdersById(ids string) (response string, err error) {
	targetUrl := BaseUrl + "api/v2/orders?range=" + ids
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}

//提交指令修改操作
func UpdateOrder(data interface{}) {
	targetUrl := BaseUrl + "api/v2/orderModifications"
	content := httputils.Post(targetUrl, data, "application/json")
	fmt.Print(content)
}

//按照id查询指令修改
func GetUpdateOrderById(id string) (response string, err error) {
	targetUrl := BaseUrl + "api/v2/orderModifications?id=" + id
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}

//系统异常码信息
func SysErrorsMessage() (response string, err error) {
	targetUrl := BaseUrl + "api/v2/errors"
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}

//任务参数信息
func TaskParameters() (response string, err error) {
	targetUrl := BaseUrl + "api/v2/taskParameters"
	resp, _ := httputils.CurlWithParam(targetUrl, "get", nil)
	fmt.Print(resp)
	return resp, nil
}
