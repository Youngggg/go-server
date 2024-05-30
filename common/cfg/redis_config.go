package cfg

import (
	"apple/common/env"
)

type RedisConfig struct {
	Address  string
	Password string
	Db       int
}

// 获取默认redis配置
func GetDefaultRedisConfig() RedisConfig {
	switch env.Get() {
	case "dev":
		return RedisConfig{
			Address:  "35.187.225.32:6379",
			Password: "aje25!*djDK2S",
			Db:       0,
		}
	case "test":
		return RedisConfig{
			Address:  "35.187.225.32:6379",
			Password: "aje25!*djDK2S",
			Db:       0,
		}
	case "prod": // todo
		return RedisConfig{
			Address:  "10.148.0.6:6379",
			Password: "aje25!*djDK2S",
			Db:       0,
		}
	default:
		panic("没有合适的profile: " + env.Get())
	}
}

// ai只部署了测试环境 因此只有测试环境配置
func GetAiRedisConfig() RedisConfig {
	switch env.Get() {
	case "dev":
		return RedisConfig{}
	case "test":
		return RedisConfig{
			Address:  "10.148.0.8:6379", // 对应筛料测试机
			Password: "aje25!*djDK2S",
			Db:       2,
		}
	case "prod":
		return RedisConfig{
			Address:  "127.0.0.1:6379", // 对应新ai
			Password: "aje25!*djDK2S",
			Db:       2,
		}
	default:
		panic("没有合适的profile: " + env.Get())
	}
}
