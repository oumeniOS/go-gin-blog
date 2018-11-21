package gredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/oumeniOS/go-gin-blog/pkg/setting"
	"fmt"
	"time"
	"encoding/json"
	)

var RedisConn *redis.Pool

func Setup() error  {
	RedisConn = &redis.Pool{
		MaxIdle:setting.RedisSetting.MaxIdle,
		MaxActive:setting.RedisSetting.MaxActive,
		IdleTimeout:setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",fmt.Sprintf("%s:%s",setting.RedisSetting.Host,setting.RedisSetting.Port))
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH",setting.RedisSetting.Password); err != nil{
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

//保存json串
func Set(key string, data interface{}, time int) error  {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil{
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool  {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func Get(key string)([]byte, error)  {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET",key))
	if err != nil{
		return nil, err
	}
	return reply, nil
}

func Delete(key string)(bool, error)  {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL",key))
}

func LikeDeletes(key string) error  {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("*%s*",key)))
	if err != nil{
		return err
	}
	for _, key := range keys{
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil

}









































