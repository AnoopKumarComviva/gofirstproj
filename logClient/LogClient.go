package logClient

import (
	"log"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Info (logStr string) {
	log.Print("[INFO] ",logStr)
}
func Trace (logStr string) {
	log.Print("[TRACE] ",logStr)
}

func Redis (logStr string) {
	var conn redis.Conn

	if conn, err := redis.Dial("tcp", "172.25.26.61:6379", redis.DialPassword("redispass")); err != nil {
		log.Fatal(err)
		conn.Close()
		return
	}
	defer conn.Close()

	_, err := conn.Do("HSET", "LOGGER", "INFO", logStr)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("**Redis Insertion Completed**")
}