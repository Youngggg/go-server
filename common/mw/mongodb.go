package mw

import (
	"context"
	"fmt"
	"time"

	"apple/common/cfg"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	logFormat     = "[mongodao] [info]  %s [CMD] %s\n"
	logTimeFormat = "2006/01/02 15:04:05.999999"
)

// 初始化mongo链接
func InitMongoDriver(addr string, mongoDBConnectionSetting cfg.MongoDbConnectionSetting) *mongo.Client {
	want, err := readpref.New(readpref.SecondaryMode) // 表示只使用辅助节点
	if err != nil {
		panic(fmt.Sprintf("init mongodb error: %s", err.Error()))
	}
	wc := writeconcern.New(writeconcern.WMajority())
	readconcern.Majority()

	// 设置选项
	opt := options.Client().ApplyURI(addr)
	opt.SetLocalThreshold(time.Duration(mongoDBConnectionSetting.LocalThreshold) * time.Second)   // 只使用与mongo操作耗时小于设置秒数的 默认15毫秒
	opt.SetMaxConnIdleTime(time.Duration(mongoDBConnectionSetting.MaxConnIdleTime) * time.Second) // 指定连接可以保持空闲的最大秒数 默认为0即不删除空闲链接
	opt.SetMaxPoolSize(uint64(mongoDBConnectionSetting.MaxPoolSize))                              // 使用最大的连接数
	opt.SetReadPreference(want)                                                                   // 表示只使用辅助节点
	opt.SetReadConcern(readconcern.Majority())                                                    // 指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
	opt.SetWriteConcern(wc)                                                                       // 请求确认写操作传播到大多数mongodb实例

	// 打印cmd
	if mongoDBConnectionSetting.ShowCmd {
		cmdMonitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				fmt.Printf(logFormat, time.Now().Format(logTimeFormat), evt.Command)
			},
		}
		opt.SetMonitor(cmdMonitor)
	}

	// 连接
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(fmt.Sprintf("init mongodb error: %s", err.Error()))
	}

	// 测试连接是否可用
	ctx, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Sprintf("init mongodb error: %s", err.Error()))
	}

	return client
}
