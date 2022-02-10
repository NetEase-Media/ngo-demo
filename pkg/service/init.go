package service

import (
	"github.com/NetEase-Media/ngo/client/db"
	"github.com/NetEase-Media/ngo/client/kafka"
	"github.com/NetEase-Media/ngo/client/redis"
)

var producer *kafka.Producer
var consumer *kafka.Consumer
var mysqlDB *db.Client
var DDB *db.Client
var redisClient redis.Redis
var shardedRedisClient redis.Redis

// Init 初始化包内变量
func Init() {
	producer = kafka.GetProducer("k1")
	consumer = kafka.GetConsumer("k1")

	mysqlDB = db.GetMysqlClient("db01")

	redisClient = redis.GetClient("redis01")
	shardedRedisClient = redis.GetClient("sharedsentinel01")

	Consume() // 启动消费者
}
