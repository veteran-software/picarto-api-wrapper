/*
 * Copyright (c) 2023. Veteran Software
 *
 * Picarto API Wrapper - A custom wrapper for the Picarto REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

var Rest *RateLimiter

type RateLimiter struct {
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

func NewPicarto(clientId string, clientSecret ...string) *RateLimiter {
	token = clientId
	if clientSecret != nil && len(clientSecret) == 1 {
		secret = &clientSecret[0]
	}

	return &RateLimiter{
		bucket: &bucket{
			Remaining: 1,
			reset:     time.Now().UTC().Truncate(60 * time.Second).Add(1 * time.Minute),
		},
	}
}

func (r *RateLimiter) getWaitTime(minRemaining int) time.Duration {
	if r.bucket.Remaining < minRemaining && r.bucket.reset.After(time.Now()) {
		return r.bucket.reset.Sub(time.Now())
	}

	return 0
}

func (r *RateLimiter) lockBucketObject() {
	r.bucket.Lock()

	if wait := r.getWaitTime(1); wait > 0 {
		time.Sleep(wait)
	}

	r.bucket.Remaining--
}

func (r *RateLimiter) lockBucket() {
	r.lockBucketObject()
}

func (b *bucket) release(headers http.Header) error {
	defer b.Unlock()

	if headers == nil {
		return nil
	}

	//goland:noinspection SpellCheckingInspection
	remaining := headers.Get("x-ratelimit-remaining")

	// Picarto does not expose a specific reset time in the headers since they reset at the top of each minute
	// So we're going to parse the Date header and do it ourselves
	serverDate := headers.Get("Date")

	if serverDate != "" {
		parsedDate, _ := time.Parse(time.RFC1123, serverDate)

		// This prevents accidental rounding up an extra minute
		var d time.Duration
		if parsedDate.Second() >= 30 {
			d = 0 * time.Minute
		} else {
			d = 1 * time.Minute
		}
		b.reset = parsedDate.Add(d).Round(1 * time.Minute)
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
