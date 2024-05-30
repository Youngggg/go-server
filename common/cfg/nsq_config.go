package cfg

import "apple/common/env"

func GetNsqConfig() string {
	switch env.Get() {
	case "dev":
		return "35.187.225.32:4150"
	case "test":
		return "35.187.225.32:4150"
	case "prod":
		return "10.148.0.6:4150"
	default:
		panic("没有合适的profile: " + env.Get())
	}
}
