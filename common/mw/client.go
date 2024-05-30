package mw

import (
	"apple/common/cfg"

	"github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

var Init = &_Init{}

var Client = &_Client{}

type _Init struct {
}

type _Client struct {
	AiRedis        *redis.Client // ai redis client
	DefaultRedis   *redis.Client
	DefaultMongodb *mongo.Client
	NsqProducer    *nsq.Producer
}

// 初始化默认redis db=0
func (*_Init) DefaultRedis() {
	Client.DefaultRedis = initRedis(cfg.GetDefaultRedisConfig())
}

func (*_Init) DefaultMongodb() {
	Client.DefaultMongodb = InitMongoDriver(cfg.DefaultMongoDbAddrConfig(), cfg.DefaultMongoDBConnectionSetting())
}
