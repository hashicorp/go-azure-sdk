// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transport

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

const maxDrainBytes = 1 << 20 // 1 MiB

// New returns a configured *http.Transport with connection pooling, timeouts,
// and HTTP/2 health-check pings enabled.
func New(useProxy bool) *http.Transport {
	var proxyFunc func(*http.Request) (*url.URL, error)
	if useProxy {
		proxyFunc = http.ProxyFromEnvironment
	}

	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	return &http.Transport{
		Proxy:       proxyFunc,
		DialContext: dialer.DialContext,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   16,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     true,
		HTTP2: &http.HTTP2Config{
			SendPingTimeout: 30 * time.Second,
			PingTimeout:     15 * time.Second,
		},
	}
}

// DrainAndClose drains up to 1 MiB from body and closes it so the underlying
// HTTP connection can be returned to the idle pool.
func DrainAndClose(body io.ReadCloser) {
	if body == nil || body == http.NoBody {
		return
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(body, maxDrainBytes))
	_ = body.Close()
}
