package service

import (
	"context"
	"time"

	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/go-redis/redis/v8"
)

// KEY 测试用的key
const KEY = "test_key"

//EXPIRE 过期时间
const EXPIRE = time.Minute * 5

// Write 写缓存
func Write(message string) error {
	_, err := redisClient.Set(context.Background(), KEY, message, EXPIRE).Result()

	if err != nil {
		return err
	}
	return nil
}

// Read 读取数据
func Read() string {
	value, err := redisClient.Get(context.Background(), KEY).Result()
	if err != nil && err != redis.Nil { // 如果key 不存在，redis 会报出异常
		log.Error(err.Error())
		return ""
	}
	return value
}

// ShardedWrite 写缓存
func ShardedWrite(message string) error {
	_, err := shardedRedisClient.Set(context.Background(), KEY, message, EXPIRE).Result()

	if err != nil {
		return err
	}
	return nil
}

// ShardedRead 读取数据
func ShardedRead() string {
	value, err := shardedRedisClient.Get(context.Background(), KEY).Result()
	if err != nil && err != redis.Nil { // 如果key 不存在，redis 会报出异常
		log.Error(err.Error())
		return ""
	}
	return value
}

// ShardedPipe 读取数据
func ShardedPipe() string {
	ctx := context.Background()
	pipe := shardedRedisClient.Pipeline()
	pipe.Set(ctx, "key1", "value1", EXPIRE)
	pipe.HGet(ctx, "key2", "field1")
	cmds, _ := pipe.Exec(ctx)
	for i := range cmds {
		log.Info(cmds[i].Name(), cmds[i].FullName(), cmds[i].Args(), cmds[i].Err())
	}

	value, err := shardedRedisClient.Get(ctx, "key1").Result()
	if err != nil && err != redis.Nil { // 如果key 不存在，redis 会报出异常
		log.Error(err.Error())
		return ""
	}
	return value
}

// 多处读取
func MultiRead() []string {
	var keys = []string{"k1", "k2"}
	ret := batchGet(keys...)
	return ret
}

func batchGet(keys ...string) []string {

	ret := redisClient.MGet(context.Background(), keys...)
	if ret == nil {
		return nil
	}
	if ret.Err() != nil && ret.Err() != redis.Nil {
		return nil
	}

	var s []string
	for _, v := range ret.Val() {
		switch t := v.(type) {
		case string:
			s = append(s, t)
		}
	}
	return s
}
