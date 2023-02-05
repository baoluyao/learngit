package cachelib

import (
	"errors"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

const Nil = "redigo: nil returned"

type CACHE struct {
	poolrw    *redis.Pool
	poolronly *redis.Pool
}

func NewCache() *CACHE {
	return &CACHE{}
}

/*
rwServer : read & write server address
rServer: read only server address
*/
func (t *CACHE) Init(rwServer string, rServer string) {
	if len(rwServer) == 0 && len(rServer) == 0 {
		panic("server empty")
	}

	if len(rwServer) > 0 {
		t.poolrw = &redis.Pool{
			MaxIdle:     10,
			IdleTimeout: 50 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", rwServer)
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}

	if len(rServer) > 0 {
		t.poolronly = &redis.Pool{
			MaxIdle:     10,
			IdleTimeout: 50 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", rServer)
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}
}

func (t *CACHE) GetRWConn() rwConn {
	if t.poolrw == nil {
		return nil
	}

	out := &Conn{}
	out.conn = t.poolrw.Get()

	return out
}

func (t *CACHE) GetRConn() rConn {
	if t.poolronly == nil {
		return t.GetRWConn()
	}

	out := &Conn{}
	out.conn = t.poolronly.Get()
	return out
}

//read only
type rConn interface {
	Close() error                                       //
	GET(key string) (interface{}, error)                //
	EXISTS(key string) (bool, error)                    //only for keys of KV type//
	TTL(key string) (int64, error)                      //only for keys of KV type//
	HGET(key string, field string) (interface{}, error) //
	HGETALL(key string) ([]string, error)               //
	HVALS(key string) ([]string, error)
	HMGET(key string, fields ...interface{}) ([]string, error) //
	HKEYS(key string) ([]string, error)                        //
	HLEN(key string) (int64, error)                            //
	HEXISTS(key string, field string) (bool, error)
	ZKEYSBYSCORE(name string, score_start int64, score_end int64, limit int64) ([]string, error)
	ZEXISTS(name string, key string) (bool, error) //ssdb only
	ZSIZE(name string) (int64, error)
	ZCOUNT(name string, score_start int64, score_end int64) (int64, error)
	ZSCORE(name string, key string) (int64, error)
	MULTI_ZGET(name string, keys ...string) (map[string]int64, error) //ssdb only
}

//write
type wConn interface {
	SET(key string, val interface{}) (string, error)                  //
	SETNX(key string, val interface{}) (bool, error)                  //
	SETEX(key string, val interface{}, ttlsecond int) (string, error) //
	DEL(key string) (bool, error)                                     //only for keys of KV type //
	EXPIRE(key string, ttlsecond int) (bool, error)                   //only for keys of KV type //
	HSET(key string, field string, val interface{}) (int64, error)    //
	HDEL(key string, fields ...interface{}) (int64, error)            //
	HMSET(key string, args ...interface{}) (string, error)            //
	HCLEAR(key string) error                                          //ssdb only
	ZADD(name string, args ...interface{}) (int64, error)
	ZREM(name string, keys ...string) (int64, error)
	ZCLEAR(name string) (int64, error) //ssdb only
}

//read & write
type rwConn interface {
	rConn
	wConn
}

type Conn struct {
	conn redis.Conn
}

func (t *Conn) Close() error {
	return t.conn.Close()
}

func (t *Conn) GET(key string) (interface{}, error) {
	return t.conn.Do("get", key)
}

func (t *Conn) EXISTS(key string) (bool, error) {
	return redis.Bool(t.conn.Do("exists", key))
}

func (t *Conn) TTL(key string) (int64, error) {
	return redis.Int64(t.conn.Do("ttl", key))
}

func (t *Conn) HGET(key string, field string) (interface{}, error) {
	return t.conn.Do("hget", key, field)
}

func (t *Conn) HGETALL(key string) ([]string, error) {
	return redis.Strings(t.conn.Do("hgetall", key))
}

func (t *Conn) HVALS(key string) ([]string, error) {
	return redis.Strings(t.conn.Do("hvals", key))
}

func (t *Conn) HMGET(key string, fields ...interface{}) ([]string, error) {
	arglist := make([]interface{}, 0, 1+len(fields))
	arglist = append(arglist, key)
	arglist = append(arglist, fields...)
	return redis.Strings(t.conn.Do("hmget", arglist...))
}

func (t *Conn) HKEYS(key string) ([]string, error) {
	return redis.Strings(t.conn.Do("hkeys", key))
}

func (t *Conn) HLEN(key string) (int64, error) {
	return redis.Int64(t.conn.Do("hlen", key))
}

func (t *Conn) HEXISTS(key string, field string) (bool, error) {
	return redis.Bool(t.conn.Do("hexists", key, field))
}

func (t *Conn) SET(key string, val interface{}) (string, error) {
	return redis.String(t.conn.Do("set", key, val))
}

//true : if key not exist and setting key success
func (t *Conn) SETNX(key string, val interface{}) (bool, error) {
	return redis.Bool(t.conn.Do("setnx", key, val))
}

func (t *Conn) SETEX(key string, val interface{}, ttlsecond int) (string, error) {
	return redis.String(t.conn.Do("setex", key, ttlsecond, val))
}

//true:  delete key success or key not exist
func (t *Conn) DEL(key string) (bool, error) {
	return redis.Bool(t.conn.Do("del", key))
}

//true :  key exist and setting ttl success
//false : if key not exist or can not set ttl
func (t *Conn) EXPIRE(key string, ttlsecond int) (bool, error) {
	return redis.Bool(t.conn.Do("expire", key, ttlsecond))
}

//1 : if field is newly created, and setting val success
//0 : if field already exist, and old val updated to new val
func (t *Conn) HSET(key string, field string, val interface{}) (int64, error) {
	return redis.Int64(t.conn.Do("hset", key, field, val))
}

//return field count been deleted. if field not exist,then it's ignored.
func (t *Conn) HDEL(key string, fields ...interface{}) (int64, error) {
	arglist := make([]interface{}, 0, 1+len(fields))
	arglist = append(arglist, key)
	arglist = append(arglist, fields...)
	return redis.Int64(t.conn.Do("hdel", arglist...))
}

func (t *Conn) HMSET(key string, args ...interface{}) (string, error) {
	arglist := make([]interface{}, 0, 1+len(args))
	arglist = append(arglist, key)
	arglist = append(arglist, args...)
	return redis.String(t.conn.Do("hmset", arglist...))
}

//delete all keys in a hashmap. delete hashmap itself
func (t *Conn) HCLEAR(key string) error {
	_, err := t.conn.Do("hclear", key)
	return err
}

//insert many key-score pairs into a Sorted Set. if key already exist, its score will be updated.
//score must be integer number(negtive number supported). MUST NOT be float number
//eg. ZADD(name, score1, key1, score2, key2, score3, key3)
//return num of keys newly inserted, not including keys already there.
func (t *Conn) ZADD(name string, args ...interface{}) (int64, error) {
	if len(args) == 0 {
		return 0, nil
	}
	arglist := make([]interface{}, 0, 1+len(args))
	arglist = append(arglist, name)
	arglist = append(arglist, args...)

	return redis.Int64(t.conn.Do("zadd", arglist...))
}

//delete key-score pairs in a Sorted Set.
//if key not there, it's ignored.
//return num of keys actually deleted.
func (t *Conn) ZREM(name string, keys ...string) (int64, error) {
	if len(keys) == 0 {
		return 0, nil
	}

	arglist := make([]interface{}, 0, 1+len(keys))
	arglist = append(arglist, name)
	for _, key := range keys {
		arglist = append(arglist, key)
	}

	return redis.Int64(t.conn.Do("zrem", arglist...))
}

//test if a key exist in a Sorted Set
func (t *Conn) ZEXISTS(name string, key string) (bool, error) {
	out, err := redis.Strings(t.conn.Do("zexists", name, key))
	if err != nil {
		return false, err
	}

	if len(out) != 1 {
		return false, errors.New("command error")
	}

	if out[0] == "1" {
		return true, nil
	}

	return false, nil
}

//enum keys in a Sorted Set
//score must be integer number(negtive number supported). MUST NOT be float number
func (t *Conn) ZKEYSBYSCORE(name string, score_start int64, score_end int64, limit int64) ([]string, error) {
	return redis.Strings(t.conn.Do("zrangebyscore", name, score_start, score_end, "limit", 0, limit))
}

//Returns the number of elements of the sorted set stored at the specified key which have scores in the range [start,end]
func (t *Conn) ZCOUNT(name string, score_start int64, score_end int64) (int64, error) {
	return redis.Int64(t.conn.Do("zcount", name, score_start, score_end))
}

//num of elements in a Sorted Set
func (t *Conn) ZSIZE(name string) (int64, error) {
	return redis.Int64(t.conn.Do("zcard", name))
}

//delete all elements in a Sorted Set, delete the Sorted Set ifself
//return num of keys deleted.
func (t *Conn) ZCLEAR(name string) (int64, error) {
	out, err := redis.Strings(t.conn.Do("zclear", name))
	if err != nil {
		return 0, err
	}

	if len(out) != 1 {
		return 0, errors.New("command error")
	}

	return strconv.ParseInt(out[0], 10, 64)
}

//get the score of a key in a Sorted Set
func (t *Conn) ZSCORE(name string, key string) (int64, error) {
	out, err := redis.String(t.conn.Do("zscore", name, key))
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(out, 10, 64)
}

//get scores of keys in a Sorted Set. if key not exist, it's ignored.
//return map[key]score
func (t *Conn) MULTI_ZGET(name string, keys ...string) (map[string]int64, error) {
	scores := make(map[string]int64)

	if len(keys) == 0 {
		return scores, nil
	}

	arglist := make([]interface{}, 0, 1+len(keys))
	arglist = append(arglist, name)
	for _, key := range keys {
		arglist = append(arglist, key)
	}

	out, err := redis.Strings(t.conn.Do("multi_zget", arglist...))
	if err != nil {
		return nil, err
	}

	if len(out)%2 != 0 {
		return nil, errors.New("command error")
	}

	for i := 0; i < len(out); i += 2 {
		nScore, err := strconv.ParseInt(out[i+1], 10, 64)
		if err != nil {
			return nil, err
		}

		scores[out[i]] = nScore
	}

	return scores, nil
}
