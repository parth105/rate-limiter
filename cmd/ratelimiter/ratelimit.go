package ratelimiter

import (
	"fmt"
	"time"
)

type RateLimit struct {
	userDB       map[string][]int64
	window       int64
	requestLimit int64
	//backendUrl		string
}

func NewRateLimit(w, r int64) *RateLimit {
	return &RateLimit{
		userDB:       map[string][]int64{},
		window:       w,
		requestLimit: r,
	}
}

func (r *RateLimit) Call(user string) (rateLimited bool) {
	now := time.Now()
	timeInt := now.Unix()

	// Returns val, true if the user is within the time window or the number of requests
	if _, i := r.userDB[user]; !i { // New user, add to DB
		r.userDB[user] = []int64{timeInt, 1}
		fmt.Printf("Call #1, for new user %s in time interval %v, logged with # of requests 1; ratelimited: ", user, timeInt)
		rateLimited = false
	} else { // Existing user, check timeinterval
		if s := r.userDB[user]; timeInt-s[0] < r.window {
			r.userDB[user][1] += 1
			if s[1] <= r.requestLimit { // Existing user, in timeinterval and request limit
				fmt.Printf("Call #%v, for existing user %s in time interval %v, and requestlimit %v; ratelimited: ", s[1], user, timeInt, r.requestLimit)
				rateLimited = false
			} else { // Existing user, in timeinterval and request limit
				fmt.Printf("Call #%v, for existing user %s in time interval %v, but outside requestlimit %v; ratelimited: ", s[1], user, timeInt, r.requestLimit)
				rateLimited = true
			}
		} else { // Existing user, quota refreshed after time interval elapsed
			r.userDB[user] = []int64{timeInt, 1}
			fmt.Printf("Call #1, for existing user %s outside time interval %v, in requestlimit %v; ratelimited: ", user, timeInt, r.requestLimit)
			rateLimited = false
		}
	}
	return // Named return rateLimited bool
}
