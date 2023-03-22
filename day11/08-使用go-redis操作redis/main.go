package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6739",
		Password: "",
		DB: 0,

	})

	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//go get github.com/go-redis/redis
	//0代表过期时间，0代表永久，设置10代表10s
	err := rdb.Set("name1", "zhangsan", 0).Err()
	if err != nil {
		panic(err)
	}

	v, err := rdb.Get("name1").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)
	
}