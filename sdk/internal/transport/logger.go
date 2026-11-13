// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transport

import (
	"io"
	"log"
	"regexp"
)

var sensitiveQueryParamRegexp = regexp.MustCompile(`(?i)([?&](?:sig|sv|se|sp|sr|code|api-key|token|client_secret)=)([^&\s"']+)`)

type redactingWriter struct {
	w io.Writer
}

func (rw redactingWriter) Write(p []byte) (int, error) {
	redacted := sensitiveQueryParamRegexp.ReplaceAll(p, []byte("${1}REDACTED"))
	if _, err := rw.w.Write(redacted); err != nil {
		return 0, err
	}
	return len(p), nil
}

// RedactingLogger returns a *log.Logger that delegates to log.Default() while
// redacting sensitive query-string parameters (such as SAS tokens and secrets).
func RedactingLogger() *log.Logger {
	d := log.Default()
	return log.New(redactingWriter{w: d.Writer()}, d.Prefix(), d.Flags())
}

// NewRedactingLogger returns a *log.Logger writing redacted output to w.
func NewRedactingLogger(w io.Writer) *log.Logger {
	return log.New(redactingWriter{w: w}, "", 0)
}
