package cfg

var TaskConfig TaskConfigDetail

// 全局配置
type TaskConfigDetail struct {
	AppConfig    AppConfig    `mapstructure:"app_config,omitempty"`    // 项目基础配置
	IpConfig     IpConfig     `mapstructure:"ip_config,omitempty"`     // ip相关配置
	LogConfig    LogConfig    `mapstructure:"log_config,omitempty"`    // 日志相关配置
	SwitchConfig SwitchConfig `mapstructure:"switch_config,omitempty"` // 开关配置
}

// gin框架配置
type AppConfig struct {
	Mode string `mapstructure:"mode,omitempty"` // gin三种模式 release、debug、test
	Port string `mapstructure:"port,omitempty"` // gin服务监听端口
}

// ip配置
type IpConfig struct {
	IpProxyProviders string `mapstructure:"ip_proxy_providers,omitempty"`
	BaseIpApiUrl     string `mapstructure:"base_ip_api_url,omitempty"` // 10086(自建ip供应商)api-地址
	PyIpApiUrl       string `mapstructure:"py_ip_api_url,omitempty"`   // py(第三方ip供应商)api-地址
}

// 开关配置
type SwitchConfig struct {
	FetchIpFromServer bool `mapstructure:"fetch_ip_from_server"` // 是否从api取ip 目前用到地方(在跑的任务ip不够时从api请求获取ip)
}

func GetTaskConfig() TaskConfigDetail {
	return TaskConfig
}
