// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import "net/http"

// DisableKeepAlives configures c with a transport that disables HTTP keep-alives
// so net/http.Transport does not auto-retry simulated EOFs on reused idle connections.
func DisableKeepAlives(c interface{ SetTransport(http.RoundTripper) }) {
	c.SetTransport(&http.Transport{DisableKeepAlives: true})
}
