package main

import (
	memorycache "myMod/Memory-cache"
	"time"
)

func main() {
	c := memorycache.Cache{}
	cache := c.New()
	cache.Set("userId", 42, time.Second*5)
	
	for i := 0; i < 6; i++ {
		userId, err := cache.Get("userId")
		if err != nil {
			println("Time is over")
			break
		}
		time.Sleep(time.Second)
		println(userId)
	}

}
