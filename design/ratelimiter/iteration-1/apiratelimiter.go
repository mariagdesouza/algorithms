package main

import (
	"fmt"
	"time"
)

/*






 */

// Token Bucket
// Fills at a pre-determined rate

type RateLimiter struct {
	capacity     int64
	leakInterval int           //in millisecondss
	tokenBucket  chan struct{} // Queue - FIFO
	quit         chan bool
}

func NewRateLimiter(c int64, leakInterval int) *RateLimiter {
	r := RateLimiter{capacity: c, leakInterval: 2}
	r.tokenBucket = make(chan struct{}, c)

	// Start the leaker.
	go func() {
		ticker := time.NewTicker(time.Duration(leakInterval) * time.Millisecond)
		for range ticker.C {
			if len(r.tokenBucket) > 0 {
				fmt.Println("leaking")
				<-r.tokenBucket
			}
		}
	}()
	return &r
}

func (r *RateLimiter) Close() {
	close(r.tokenBucket)
}

func (r *RateLimiter) Hit() bool {
	if int64(len(r.tokenBucket)) < r.capacity {
		r.tokenBucket <- struct{}{}
		return true
	}
	select {
	case <-time.After(time.Duration(r.leakInterval) * time.Millisecond):
		return false
	}

}

func main() {

	r := NewRateLimiter(5, 1)
	for i := 0; i < 20; i++ {
		if r.Hit() {
			fmt.Println("Hit Successful:", i)
		} else {
			fmt.Println("Hit Failed:", i)
		}
	}
	r.Close()
}
