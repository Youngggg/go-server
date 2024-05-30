package boluo

import (
	"fmt"
	"testing"
)

func Test_QueryUserInfo(t *testing.T) {

	controller := NewApiBoluoController()
	result := controller.QueryUserInfo()
	fmt.Println(result)
}

func Test_QueryTgCountry(t *testing.T) {

	controller := NewApiBoluoController()
	result := controller.QueryTgCountry()
	fmt.Println(result)
}

func Test_QueryTgTaskInfo(t *testing.T) {

	taskId := "28878"
	controller := NewApiBoluoController()
	result := controller.QueryTgTaskInfo(taskId)
	fmt.Println(result)
}

func Test_AddTgTask(t *testing.T) {

	controller := NewApiBoluoController()
	result := controller.AddTgTask("印度", "E:\\test_ID_1000.txt")

	if result != nil {
		fmt.Println(result)
		fmt.Println(result.Data.TaskId)
	} else {
		fmt.Println("result nil")
	}
}

func Test_DownloadTgTask(t *testing.T) {

	taskId := "28878"
	controller := NewApiBoluoController()
	result := controller.DownloadTgTask(taskId)
	fmt.Println(result)
	//&{1 success 1716026123 {http://api.shaixuan.vip/tgopen/2024-05-28/28878_全部数据_217个.csv}}
}
