package ratelimiter

import (
	"time"
)

type RateLimit struct {
	userDB       map[string]map[int64]int64
	window       int64
	requestLimit int64
	//backendUrl		string
}

func NewRateLimit(w int64, r int64) *RateLimit {
	return &RateLimit{
		userDB:       map[string]map[int64]int64{},
		window:       w,
		requestLimit: r,
	}
}

func (r *RateLimit) Call(user string) bool {
	now := time.Now()
	timeInt := now.Unix()

	// Return true if the user is within the time window or the number of requests
	if _, i := r.userDB[user]; !i {
		r.userDB[user] = map[int64]int64{}
		r.userDB[user][timeInt] = 1
		//fmt.Printf("Call #1, for new user %s in time interval %v, logged with # of requests 1; returning ", user, timeInt)
		return true
	} else {
		if _, i := r.userDB[user][timeInt]; !i {
			r.userDB[user][timeInt] = 1
			//fmt.Printf("Call #1, for existing user %s in time interval %v, logged with # of requests 1; returning ", user, timeInt)
			return true
		} else {
			r.userDB[user][timeInt] += 1
			//fmt.Printf("Call #%v, for existing user %s in time interval %v; returning ", r.userDB[user][timeInt], user, timeInt)
			if r.userDB[user][timeInt] > r.requestLimit {
				return false
			} else {
				return true
			}
		}
	}
}
