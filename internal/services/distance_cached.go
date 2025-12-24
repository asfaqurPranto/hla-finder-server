package services

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		// Optional: Password: "", DB: 0
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Redis connection failed: %v", err))
	}
}

func SetDistance(source,destination string ,dist int) error{

	rdb:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	//
	//key=source ,dest, val
	err:=rdb.HSet(ctx,source,destination,dist).Err()
	if err!=nil{
		return err
	}
	err=rdb.Expire(ctx, source, 30*24*time.Hour).Err()
	if err!=nil{
		return err
	}
	//reverse mapping korlam
	err=rdb.HSet(ctx,destination,source,dist).Err()
	if err!=nil{
		return err
	}
	err=rdb.Expire(ctx, destination, 30*24*time.Hour).Err()
	if err!=nil{
		return err
	}
	fmt.Println("cached done for: ",source,destination,dist)
	return nil
}

func GetDistance(source, destination string) (int, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	valStr,err:= rdb.HGet(ctx, source, destination).Result()
	if err!=nil{
		return -1,err
	}
	valInt, err := strconv.Atoi(valStr)
	if err!=nil{
		return -1,err
	}
	return valInt,nil

}

