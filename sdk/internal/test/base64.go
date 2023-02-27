// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import (
	"encoding/base64"
	"testing"
)

func Base64DecodeCertificate(t *testing.T, clientCertificate string) (pfx []byte) {
	if clientCertificate != "" {
		out := make([]byte, base64.StdEncoding.DecodedLen(len(clientCertificate)))
		n, err := base64.StdEncoding.Decode(out, []byte(clientCertificate))
		if err != nil {
			t.Fatalf("test.NewConnection(): could not decode value of CLIENT_CERTIFICATE: %v", err)
		}
		pfx = out[:n]
	}
	return
}
