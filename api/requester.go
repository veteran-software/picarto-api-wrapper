package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
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

func (r *rateLimiter) Request(method, route string, data *interface{}, token interface{}) (*http.Response, error) {
	return r.requestWithLockedBucket(method, route, "application/json", data)
}

func (r *rateLimiter) requestWithLockedBucket(method, route, contentType string, b *interface{}) (*http.Response, error) {
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
		log.Warnln("Rate Limited!")
		log.Infoln(route)
		log.Infoln(resp.Status)

		time.Sleep(time.Until(r.bucket.reset))

		resp, err = r.requestWithLockedBucket(method, route, contentType, b)
	}

	return resp, nil
}
