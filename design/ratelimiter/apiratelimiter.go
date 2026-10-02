package main

import (
	"fmt"
	"time"
)

/*

LEaky Bucket Algorithm
Allows hits till a capacity is hit and

Improvement

GCRA
Generic Cell Rate Algo


*/

// Token Bucket
// Fills at a pre-determined rate

const (
	SLEEPMICRO = 100
)

type RateLimiter struct {
	capacity     int64
	leakInterval int           // in millisecondss
	tokenBucket  chan struct{} // bucket - leaks when full
	sleepFor     time.Duration
}

func NewRateLimiter(c int64, leakInterval int) *RateLimiter {
	//create the bucket
	r := RateLimiter{capacity: c, leakInterval: 30}
	r.tokenBucket = make(chan struct{}, c)
	r.sleepFor = time.Duration(SLEEPMICRO * time.Microsecond)
	// Start the leaker.
	go func() {
		for {
			select {
			case <-time.After(time.Duration(r.leakInterval) * time.Microsecond):
				if len(r.tokenBucket) > 0 {
					<-r.tokenBucket
				}
			}

		}
	}()
	return &r
}

func (r *RateLimiter) Hit() bool {
	if int64(len(r.tokenBucket)) < r.capacity {
		r.tokenBucket <- struct{}{}
		return true
	}
	return r.Sleep(r.sleepFor)
}

func (r *RateLimiter) Sleep(sleep time.Duration) bool {
	select {
	case <-time.After(sleep):
		return false
	}
}

func (r *RateLimiter) Close() {
	close(r.tokenBucket)
}

func main() {

	r := NewRateLimiter(10, 1)
	var passed, failed int
	for i := 0; i < 50; i++ {
		if r.Hit() {
			//fmt.Println("Hit Successful:", i)
			passed++
		} else {
			failed++
			r.Sleep(time.Duration(10 * time.Millisecond))
			//fmt.Println("Hit Failed:", i)
		}
	}
	r.Close()
	fmt.Println("Successful hits:", passed, " Hits Failed:", failed)
}
