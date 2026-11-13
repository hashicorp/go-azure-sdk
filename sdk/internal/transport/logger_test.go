// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transport

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewRedactingLogger(t *testing.T) {
	var buf bytes.Buffer
	l := NewRedactingLogger(&buf)

	l.Printf("[DEBUG] GET https://acct.blob.core.windows.net/container?sv=2023-01-01&sig=supersecret123&restype=container")
	got := buf.String()
	if strings.Contains(got, "supersecret123") {
		t.Fatalf("expected SAS sig to be redacted, got: %s", got)
	}
	if !strings.Contains(got, "sig=REDACTED") {
		t.Fatalf("expected sig=REDACTED in output, got: %s", got)
	}
	if !strings.Contains(got, "restype=container") {
		t.Fatalf("expected non-sensitive query params to remain, got: %s", got)
	}
}
