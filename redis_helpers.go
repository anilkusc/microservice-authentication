package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func RedisGet(key string) (string, error) {
	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0), redis.DialPassword(""))
	if err != nil {
		//log.Println(err)
		return "", err
	}
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		//log.Println(err)
		return value, err
	}

	return value, nil
}

func RedisSet(key string, value string) error {
	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0), redis.DialPassword(""))
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	conn.Do("SET", key, value)
	return nil
}

func RedisSetEx(key string, value string, expire string) error {
	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0), redis.DialPassword(""))
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	conn.Do("SETEX", key, expire, value)
	return nil
}

func RedisDelete(key string) {
	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0), redis.DialPassword(""))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	conn.Do("DEL", key)
}
