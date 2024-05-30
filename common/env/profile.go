package env

import (
	"flag"
	"testing"
)

var profile string

func init() {

	// 默认运行profile是dev
	profile = "dev"
	// 在flag前面执行，这样才能运行单元测试
	testing.Init()
	commandLineProfile := flag.String("profile", "", "运行profile")
	flag.Parse()

	if *commandLineProfile != "" {
		// 从命令行可以获取到profile就优先使用命令行变量
		profile = *commandLineProfile
	}
	//else if utils.Str.IsNotBlank(os.Getenv("GoEnv")) {
	//	// 如果没有命令行就从环境变量获取
	//	profile = os.Getenv("GoEnv")
	//}

	println("当前运行profile:" + profile)
}

func Get() string {
	return profile
}

func Set(profileInput string) {
	profile = profileInput
}

func IsTest() bool {
	return profile == "test"
}

func IsProd() bool {
	return profile == "prod"
}
