package boluo

/*
	{
	  "code": 0,
	  "msg": "string",
	  "time": "string",
	  "data": {
	    "id": 0,
	    "title": "string",
	    "sum_count": 0,
	    "success_count": 0,
	    "effective_count": 0,
	    "status": "string"
	  }
	}
*/
type CheckTgTaskResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Id             int    `json:"id"`
		Title          string `json:"title"`
		SumCount       int    `json:"sum_count"`
		SuccessCount   int    `json:"success_count"`
		EffectiveCount int    `json:"effective_count"`
		Status         string `json:"status"`
	} `json:"data"`
}

/*
	{
	  "code": 1,
	  "msg": "success",
	  "time": "1716022322",
	  "data": {
	    "id": 1206,
	    "nickname": "",
	    "mobile": "",
	    "avatar": "",
	    "money": "0.000000",
	    "score": "72.000000",
	    "is_agent": 0,
	    "wallet_address": "",
	    "loginip": "23.158.104.254",
	    "logintime": "2024-05-16 14:54:57"
	  }
	}
*/
type UserInfoResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Id            int    `json:"id"`
		Nickname      string `json:"nickname"`
		Mobile        string `json:"mobile"`
		Avatar        string `json:"avatar"`
		Money         string `json:"money"`
		Score         string `json:"score"`
		IsAgent       int    `json:"is_agent"`
		WalletAddress string `json:"wallet_address"`
		Loginip       string `json:"loginip"`
		Logintime     string `json:"logintime"`
	} `json:"data"`
}

/*
	{
	  "code": 1,
	  "msg": "success",
	  "time": "1679023208",
	  "data": [
	    "阿布哈兹",
	    "阿富汗",
	    "阿尔巴尼亚",
		]
	}
*/
type QueryCountryResult struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Time string   `json:"time"`
	Data []string `json:"data"`
}

/*
	{
	  "code": 1,
	  "msg": "添加成功",
	  "time": "1678874654",
	  "data": {
	    "task_id": "45"
	  }
	}
*/
type AddTgTaskResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
}

/*
	{
	  "code": 1,
	  "msg": "success",
	  "time": "1678950975",
	  "data": {
	    "url": "https://back.mangguotg.me/download/2023-03-16/2614_全部数据_81334个.csv"
	  }
	}
*/
type DownloadTgTaskResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}
