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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	log "github.com/veteran-software/nowlive-logging"
)

var (
	initialTimeout        = 5 * time.Millisecond
	maxTimeout            = 500 * time.Millisecond
	exponentFactor        = 2.0
	maximumJitterInterval = 2 * time.Millisecond

	backoff = heimdall.NewExponentialBackoff(initialTimeout, maxTimeout, exponentFactor, maximumJitterInterval)

	// Create a new retry mechanism with the backoff
	retrier = heimdall.NewRetrier(backoff)

	timeout = 1000 * time.Millisecond

	httpClient = httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(4),
	)
)

func (r *RateLimiter) Request(method, route string, data *interface{}) (*http.Response, error) {
	return r.requestWithLockedBucket(method, route, "application/json", data)
}

func (r *RateLimiter) requestWithLockedBucket(method, route, contentType string, b *interface{}) (*http.Response,
	error) {
	r.lockBucket()

	var buffer bytes.Buffer
	if b != nil {
		var buffer bytes.Buffer

		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&b)
		if err != nil {
			_ = r.bucket.release(nil)
			return nil, err
		}
	}

	req, err := http.NewRequest(method, route, bytes.NewReader(buffer.Bytes()))
	if err != nil {
		_ = r.bucket.release(nil)
		return nil, err
	}

	if secret != nil {
		req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bearer %s", *secret))
	}
	req.Header.Set("Client-ID", fmt.Sprintf("%s", token))
	req.Header.Set(http.CanonicalHeaderKey("Content-Type"), contentType)

	resp, err := httpClient.Do(req)
	if err != nil {
		_ = r.bucket.release(nil)
		return nil, err
	}

	err = r.bucket.release(resp.Header)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusBadGateway:
	case http.StatusTooManyRequests:
		log.Warnln(log.Picarto, log.FuncName(), "Rate Limited!")
		log.Infoln(log.Picarto, log.FuncName(), route)
		log.Infoln(log.Picarto, log.FuncName(), resp.Status)

		time.Sleep(time.Until(r.bucket.reset))

		resp, err = r.requestWithLockedBucket(method, route, contentType, b)
	}

	return resp, nil
}
