package boluo

import (
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"apple/common/utils"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type ApiBoluoController struct {
	proxyAddr string
	baseUrl   string
	apiKey    string
}

func NewApiBoluoController() *ApiBoluoController {
	b := &ApiBoluoController{
		proxyAddr: "127.0.0.1:7890",
		baseUrl:   "http://api.shaixuan.vip/api/open",
		apiKey:    "f886e417ec89e1434437ded2044a5322",
	}
	return b
}

func (self *ApiBoluoController) NewRequest() *resty.Request {
	timeOut := 30
	client := resty.New()
	client.SetTimeout(time.Duration(timeOut) * time.Second)

	if self.proxyAddr != "" {
		transport := utils.Http.BuildCommonHttpTransportWithProxyWithGoogleCipherSuites(timeOut, self.proxyAddr)
		client.SetTransport(transport)
	}

	req := client.R().SetQueryParam("api_key", self.apiKey)
	return req
}

func (self *ApiBoluoController) QueryUserInfo() *UserInfoResult {
	req := self.NewRequest()
	var result UserInfoResult
	resp, err := req.
		SetResult(&result).
		Get(self.baseUrl + "/userInfo")

	if err != nil {
		fmt.Println("ResultUserInfo.err: "+err.Error(), zap.String("resp", resp.String()))
		return nil
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("ResultUserInfo.err: ", zap.String("status", resp.Status()), zap.String("resp", resp.String()))
		return nil
	}
	return &result
}
func (self *ApiBoluoController) QueryTgCountry() *QueryCountryResult {
	req := self.NewRequest()
	var result QueryCountryResult
	resp, err := req.
		SetResult(&result).
		Get(self.baseUrl + "/tgCountry")

	if err != nil {
		fmt.Println("QueryTgCountry.err: "+err.Error(), zap.String("resp", resp.String()))
		return nil
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("QueryTgCountry.err: ", zap.String("status", resp.Status()), zap.String("resp", resp.String()))
		return nil
	}
	return &result
}
func (self *ApiBoluoController) QueryTgTaskInfo(taskId string) *CheckTgTaskResult {
	req := self.NewRequest()
	var result CheckTgTaskResult
	resp, err := req.
		SetQueryParam("task_id", taskId).
		SetResult(&result).
		Get(self.baseUrl + "/checkTgTask")

	if err != nil {
		fmt.Println("QueryTgTaskInfo.err: "+err.Error(), zap.String("resp", resp.String()))
		return nil
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("QueryTgTaskInfo.err: ", zap.String("status", resp.Status()), zap.String("resp", resp.String()))
		return nil
	}
	return &result
}

func (self *ApiBoluoController) AddTgTask(country, filePath string) *AddTgTaskResult {
	req := self.NewRequest()

	txt := utils.File.ReadWithIOUtil(filePath)
	_, fileName := path.Split(filePath)

	var result AddTgTaskResult
	resp, err := req.
		SetQueryParam("title", "tg_"+country).
		SetQueryParam("country", country).
		SetQueryParam("type", "1").
		SetMultipartField("file", fileName, "application/txt", strings.NewReader(string(txt))).
		SetResult(&result).
		Post(self.baseUrl + "/addTgTask")

	if err != nil {
		fmt.Println("AddTgTask.err: "+err.Error(), zap.String("resp", resp.String()))
		return nil
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("AddTgTask.err: ", zap.String("status", resp.Status()), zap.String("resp", resp.String()))
		return nil
	}
	return &result
}
func (self *ApiBoluoController) DownloadTgTask(taskId string) *DownloadTgTaskResult {
	req := self.NewRequest()
	var result DownloadTgTaskResult
	resp, err := req.
		SetQueryParam("task_id", taskId).
		SetResult(&result).
		Get(self.baseUrl + "/downloadTgTask")

	if err != nil {
		fmt.Println("DownloadTgTask.err: "+err.Error(), zap.String("resp", resp.String()))
		return nil
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("DownloadTgTask.err: ", zap.String("status", resp.Status()), zap.String("resp", resp.String()))
		return nil
	}
	return &result
}
