package tools

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)

var (
	RedisNoConnErr       = errors.New("can't get a redis conn")
	RedisKetamaBase      = 255
	RedisSource          map[string]string
	RedisStorageInstance *RedisStorage
)

type RedisStorage struct {
	pool map[string]*redis.Pool
	ring *HashRing
}

func newRedisStorage() *RedisStorage {
	var (
		err error
		w   int
		nw  []string
	)
	redisPool := map[string]*redis.Pool{}
	ring := NewRing(RedisKetamaBase)
	for n, addr := range RedisSource {
		nw = strings.Split(n, ":")
		if len(nw) != 2 {
			err = errors.New("node config error, it's nodeN:W")
			panic(err)
		}
		w, err = strconv.Atoi(nw[1])
		if err != nil {
			panic(err)
		}
		tmp := addr
		// WARN: closures use
		redisPool[nw[0]] = &redis.Pool{
			MaxIdle:     50,
			MaxActive:   1000,
			IdleTimeout: 28800 * time.Second,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", tmp)
				if err != nil {
					return nil, err
				}
				return conn, err
			},
		}
		ring.AddNode(nw[0], w)
	}
	ring.Bake()
	s := &RedisStorage{pool: redisPool, ring: ring}
	return s
}

// getConn get the connection of matching with key using ketama hashing.
func (s *RedisStorage) GetConn(key string) redis.Conn {
	if len(s.pool) == 0 {
		return nil
	}
	node := s.ring.Hash(key)
	p, ok := s.pool[node]
	if !ok {
		return nil
	}
	return p.Get()
}

func (s *RedisStorage) SetExpireKey(key, value string, tokenExpireTIme int) (bool, error) {
	conn := s.GetConn(key)
	defer conn.Close()
	err := conn.Send("SET", key, value)
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return false, err
	}
	err = conn.Send("EXPIRE", key, tokenExpireTIme)
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return false, err
	}
	err = conn.Flush()
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return false, err
	}
	return true, nil
}

func (s *RedisStorage) SetKey(key, value string) (bool, error) {
	conn := s.GetConn(key)
	defer conn.Close()
	err := conn.Send("SET", key, value)
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return false, err
	}
	return true, nil
}

func (s *RedisStorage) GetKey(key string) (string, error) {
	conn := s.GetConn(key)
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
	if len(value) == 0 || err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return value, nil
}

func (s *RedisStorage) ExpireKey(key string, tokenExpireTIme int) (bool, error) {
	conn := s.GetConn(key)
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, tokenExpireTIme)
	if err != nil {
		fmt.Println("redis操作失败-" + err.Error())
		return false, err
	}
	return true, nil

}

func (s *RedisStorage) DelKey(key string) (bool, error) {
	conn := s.GetConn(key)
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	if err != nil {
		fmt.Println("redis操作失败-" + err.Error())
		return false, err
	}
	return true, nil
}

func InitRedis() {
	RedisSource = make(map[string]string)
	RedisSource["node1:1"] = beego.AppConfig.String("redis_resource")
	RedisStorageInstance = newRedisStorage()
}
