package cachelib

import (
	"fmt"
	"noticeservice/global"

	"github.com/garyburd/redigo/redis"
)

var (
	cachePool *CACHE
)

func CacheInit() {
	cachePool = NewCache()
	//cachePool.Init(viper.GetString("ssdb.addr"), viper.GetString("ssdb.addr1"))
	cachePool.Init(global.SSDBSetting.Addr, global.SSDBSetting.Addr1)
}

// 获取用户key
func GetUserKey(appId uint32, userId string) (key string) {
	key = fmt.Sprintf("%d_%s", appId, userId)

	return
}

func GET(key string) (string, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return redis.String(conn.GET(key))
}

func HGETALL(key string) ([]string, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.HGETALL(key)
}

func HKEYS(key string) ([]string, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.HKEYS(key)
}

func SETNX(key string, value string) (bool, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.SETNX(key, value)
}

func SETEX(key string, value string, exp int) (string, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.SETEX(key, value, exp)
}

func SET(key string, value string) (string, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.SET(key, value)
}

func DEL(key string) (bool, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.DEL(key)
}

func HSET(key string, field string, value interface{}) (int64, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.HSET(key, field, value)
}

func HGET(key string, field string) (interface{}, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.HGET(key, field)
}

func HDEL(key string, fields ...interface{}) (int64, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.HDEL(key, fields...)
}

func HEXISTS(key string, field string) (bool, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.HEXISTS(key, field)
}

func EXPIRE(key string, exp int) (bool, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.EXPIRE(key, exp)
}

func HCLEAR(key string) error {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.HCLEAR(key)
}

func ZADD(name string, args ...interface{}) (int64, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.ZADD(name, args...)
}

func ZREM(name string, key ...string) (int64, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.ZREM(name, key...)
}

func ZKEYSBYSCORE(key string, socre_start int64, socre_end int64, limit int64) ([]string, error) {
	conn := cachePool.GetRWConn()
	defer conn.Close()
	return conn.ZKEYSBYSCORE(key, socre_start, socre_end, limit)
}
