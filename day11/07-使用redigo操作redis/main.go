package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	//go get github.com/gomodule/redigo/redis
	
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("redis.Dial() error :", err)
		return
	}

	//关闭连接
	defer conn.Close()

	_, err = conn.Do("set", "name", "jack")
	if err != nil {
		fmt.Println(err)
		return
	}

	ret, err := conn.Do("get", "name")
	if err != nil {
		fmt.Println(err)
		return
	}

	v, _ := redis.String(ret, err)
	fmt.Println("ret:", v)

}