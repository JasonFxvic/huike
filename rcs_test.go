package func_test

import (
	"testing"

	rcsinterface "github.com/JasonFxvic/huike/rcsInterface"
)

// var param map[string]string
// param := make(map[string]string)

//查询单台车辆信息
func TestGetAgvMessage(t *testing.T) {
	t.Log("测试==>查询单台车辆信息")
	var id string = "0001"
	rcsinterface.GetAgvMessage(id)
}

//获取所有机器人信息
func TestGetAllAgvMessage(t *testing.T) {
	t.Log("测试==>获取所有机器人信息")
	rcsinterface.GetAllAgvMessage()
}

//获取所有车辆正在执行的指令
func TestGetAgvOrders(t *testing.T) {
	t.Log("测试==>获取所有车辆正在执行的指令")
	id := "0001"
	rcsinterface.GetAgvOrders(id)
}

//修改车辆参数
func TestUpdateAgvAttr(t *testing.T) {
	t.Log("测试==>修改车辆参数")
	id := "0001"
	data := "{\"velocityLimit\":\"2.0\",\"powerOff\":\"false\",\"workState\":\"0\"}"
	rcsinterface.UpdateAgvAttr(id, data)
}

//创建指令
func TestCreateOrder(t *testing.T) {
	t.Log("测试==>创建指令")
	param := make(map[string]interface{})
	param["id"] = "T2102011365639"
	param["systemId"] = "WMS"
	param["type"] = "LoadingAndUnloading"
	param["source"] = "100"
	param["destination"] = "126"
	rcsinterface.CreateOrder(param)
}

//按照id查询指令
func TestGetOrderById(t *testing.T) {
	t.Log("测试==>按照id查询指令")
	var id string = "T2102011365633"
	rcsinterface.GetOrderById(id)
}

//按照id查询多个指令     ==>会查询出所有的数据
func TestGetOrdersById(t *testing.T) {
	t.Log("测试==>按照id查询多个指令")
	var ids string = "T2102011365633"
	rcsinterface.GetOrdersById(ids)
}

//提交指令修改操作
func TestUpdateOrder(t *testing.T) {
	t.Log("测试==>提交指令修改操作")
	param := make(map[string]interface{}, 11)
	param["id"] = "T2102011365635"
	param["orderId"] = "T2102011365635"
	param["type"] = 2
	param["orderPriority"] = 100

	rcsinterface.UpdateOrder(param)
}

//按照id查询指令修改
func TestGetUpdateOrderById(t *testing.T) {
	t.Log("测试==>按照id查询指令修改")
	id := "T2102011365636"
	rcsinterface.GetUpdateOrderById(id)
}

// 系统异常码信息
func TestSysErrorsMessage(t *testing.T) {
	t.Log("测试==>按照id查询指令修改")
	rcsinterface.SysErrorsMessage()
}

// 任务参数信息
func TestTaskParameters(t *testing.T) {
	t.Log("测试==>任务参数信息")
	rcsinterface.TaskParameters()
}
