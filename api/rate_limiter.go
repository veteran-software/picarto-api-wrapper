package api

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

var Rest *rateLimiter

type rateLimiter struct {
	sync.Mutex

	bucket *bucket
}

type bucket struct {
	sync.Mutex

	Remaining int
	limit     int
	reset     time.Time
	lastReset time.Time
}

func NewPicarto(clientId string, clientSecret ...string) *rateLimiter {
	token = clientId
	if clientSecret != nil && len(clientSecret) == 1 {
		secret = &clientSecret[0]
	}

	return &rateLimiter{
		bucket: &bucket{
			Remaining: 1,
			reset:     time.Now().UTC().Truncate(60 * time.Second).Add(1 * time.Minute),
		},
	}
}

func (r *rateLimiter) getWaitTime(minRemaining int) time.Duration {
	if r.bucket.Remaining < minRemaining && r.bucket.reset.After(time.Now()) {
		return r.bucket.reset.Sub(time.Now())
	}

	return 0
}

func (r *rateLimiter) lockBucketObject() {
	r.bucket.Lock()

	if wait := r.getWaitTime(1); wait > 0 {
		time.Sleep(wait)
	}

	r.bucket.Remaining--
}

func (r *rateLimiter) lockBucket() {
	r.lockBucketObject()
}

func (b *bucket) release(headers http.Header) error {
	defer b.Unlock()

	if headers == nil {
		return nil
	}

	//goland:noinspection SpellCheckingInspection
	remaining := headers.Get("x-ratelimit-remaining")
	date := headers.Get("Date")

	if date != "" {
		parsedDate, err := time.Parse(time.RFC1123, date)
		if err != nil {
			return err
		}

		b.reset = time.Unix(parsedDate.Unix(), 0)
	}

	if remaining != "" {
		parsedRemaining, err := strconv.ParseInt(remaining, 10, 32)
		if err != nil {
			return err
		}

		b.Remaining = int(parsedRemaining)
	}

	return nil
}
