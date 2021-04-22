package main

import (
	"fmt"
	"time"

	"github.com/parth105/rate-limiter/cmd/ratelimiter"
)

func main() {
	r := ratelimiter.NewRateLimit(1, 5)
	interval, _ := time.ParseDuration("0.1s")
	for i := 1; i <= 10; i++ {
		fmt.Println(r.Call("parth"))
		time.Sleep(interval)
	}

}
