package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB

	})

	ctx := context.Background()

	err := rdb.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	ress, err := rdb.ZAdd(ctx, "waitinglist:football", redis.Z{
		Score:  float64(time.Now().UnixMicro()),
		Member: 1,
	}).Result()

	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("res:", ress)
	}

}
