// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keys

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
)

func TestUpdateKeyRotationPolicyConnectionRetry(t *testing.T) {
	tests := []struct {
		name            string
		firstStatus     int
		firstError      error
		connectionFails bool
		createKey       bool
		wantAttempts    int
		wantError       bool
	}{
		{name: "connection reset during policy update", firstError: syscall.ECONNRESET, wantAttempts: 2},
		{name: "connection closed before policy response", firstError: io.EOF, wantAttempts: 2},
		{name: "connection keeps resetting", firstError: syscall.ECONNRESET, connectionFails: true, wantAttempts: 2, wantError: true},
		{name: "forbidden policy update", firstStatus: http.StatusForbidden, wantAttempts: 1, wantError: true},
		{name: "untrusted certificate", firstError: &tls.CertificateVerificationError{Err: x509.UnknownAuthorityError{}}, wantAttempts: 1, wantError: true},
		{name: "connection reset during key creation", firstError: syscall.ECONNRESET, createKey: true, wantAttempts: 1, wantError: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			const endpoint = "https://example.vault.azure.net"
			id := NewKeyID(endpoint, "example-key")
			client, err := NewKeysClientWithBaseURI(endpoint)
			if err != nil {
				t.Fatal(err)
			}
			client.Client.AuthorizeRequest = nil
			var savedURL, savedBody string
			attempts := 0
			client.Client.SetTransport(rotationPolicyRoundTripper(func(req *http.Request) (*http.Response, error) {
				attempts++
				if !test.createKey {
					if req.Method != http.MethodPut || req.URL.Host != "example.vault.azure.net" || strings.TrimLeft(req.URL.Path, "/") != "keys/example-key/rotationpolicy" {
						t.Fatalf("unexpected policy request %s %s", req.Method, req.URL)
					}
					body, err := io.ReadAll(req.Body)
					if err != nil {
						t.Fatal(err)
					}
					if attempts == 1 {
						savedURL, savedBody = req.URL.String(), string(body)
					} else if req.URL.String() != savedURL || string(body) != savedBody {
						t.Fatalf("retry sent a different URL or policy payload")
					}
				}
				status := http.StatusOK
				if test.firstError != nil && (attempts == 1 || test.connectionFails) {
					return nil, test.firstError
				}
				if attempts == 1 {
					status = test.firstStatus
				}
				return &http.Response{
					StatusCode: status,
					Header:     http.Header{"Content-Type": {"application/json"}},
					Body:       io.NopCloser(strings.NewReader(`{"attributes":{"expiryTime":"P6M"}}`)),
					Request:    req,
				}, nil
			}))
			ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
			defer cancel()
			if test.createKey {
				_, err = client.CreateKey(ctx, id, KeyCreateParameters{Kty: JsonWebKeyTypeRSA})
			} else {
				_, err = client.UpdateKeyRotationPolicy(ctx, id, KeyRotationPolicy{
					Attributes: &KeyRotationPolicyAttributes{ExpiryTime: pointer.To("P6M")},
				})
			}
			if (err != nil) != test.wantError || attempts != test.wantAttempts {
				t.Fatalf("got %d attempts and error %v, want %d attempts and error=%t", attempts, err, test.wantAttempts, test.wantError)
			}
		})
	}
}

type rotationPolicyRoundTripper func(*http.Request) (*http.Response, error)

func (f rotationPolicyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
