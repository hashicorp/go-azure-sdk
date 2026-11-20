// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transport

import (
	"bytes"
	"log"
	"strings"
	"sync"
	"testing"
)

func TestRedactingLogger(t *testing.T) {
	origOut := log.Writer()
	defer log.SetOutput(origOut)

	var buf bytes.Buffer
	log.SetOutput(&buf)

	l := RedactingLogger()
	l.Printf("[DEBUG] GET https://acct.blob.core.windows.net/container?sv=2023-01-01&sig=supersecret123&client_secret=hunter2&restype=container")

	got := buf.String()
	if strings.Contains(got, "supersecret123") || strings.Contains(got, "hunter2") {
		t.Fatalf("expected secrets to be redacted, got: %s", got)
	}
	if !strings.Contains(got, "sig=REDACTED") || !strings.Contains(got, "client_secret=REDACTED") {
		t.Fatalf("expected REDACTED placeholders in output, got: %s", got)
	}
	if !strings.Contains(got, "sv=2023-01-01") || !strings.Contains(got, "restype=container") {
		t.Fatalf("expected non-secret SAS metadata and query params to remain, got: %s", got)
	}
}

func TestRedactingLogger_Concurrent(t *testing.T) {
	origOut := log.Writer()
	defer log.SetOutput(origOut)

	var buf bytes.Buffer
	log.SetOutput(&buf)

	l := RedactingLogger()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				l.Printf("[DEBUG] GET https://login.microsoftonline.com/token?client_secret=hunter2&sig=abc")
			}
		}()
	}
	wg.Wait()
}
