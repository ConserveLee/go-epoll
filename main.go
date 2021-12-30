package main

import (
	redis "github.com/go-redis/redis/v8"
	"go-epoll/myHttp"
)

func main() {
	Foo3(myHttp.NewContext())
}

func Foo3(ctx *myHttp.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb.Set(ctx.Ctx, "key", "value", 0).Err()
}
