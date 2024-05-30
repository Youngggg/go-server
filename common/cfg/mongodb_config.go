package cfg

import (
	"apple/common/env"
)

// mongodb数据库连接设置对象
type MongoDbConnectionSetting struct {
	MaxPoolSize     int  `yaml:"max_pool_size"`
	LocalThreshold  int  `yaml:"local_threshold"`
	MaxConnIdleTime int  `yaml:"max_conn_idle_time"`
	ShowCmd         bool `yaml:"show_cmd"`
}

// mongodb数据库连接默认配置
func DefaultMongoDBConnectionSetting() MongoDbConnectionSetting {
	return MongoDbConnectionSetting{
		MaxPoolSize:     200,
		LocalThreshold:  10,
		MaxConnIdleTime: 10,
		ShowCmd:         false,
	}
}

// mongodb链接配置
func DefaultMongoDbAddrConfig() string {
	var addr string
	switch env.Get() {
	case "dev":
		addr = "mongodb://root:Xa%40503!LoginInfo@34.126.68.10:27017/?authMechanism=SCRAM-SHA-1"
	case "test":
		addr = "mongodb://root:Xa%40503!LoginInfo@10.148.0.22:27017/?authMechanism=SCRAM-SHA-1"
	case "prod":
		addr = "mongodb://root:Xa%40503!LoginInfo@10.148.0.22:27017/?authMechanism=SCRAM-SHA-1"
	default:
		panic("没有合适的profile: " + env.Get())
	}

	return addr
}
