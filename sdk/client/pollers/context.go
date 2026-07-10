// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pollers

import (
	"context"
	"time"
)

type pollKey int

const (
	skipPollingDelayKey pollKey = iota
	initialPollingDelayKey
)

// WithSkipPollingDelay returns a new context with the skip polling delay flag set.
// This is used to signal to PollUntilDone that it should not wait between polling attempts.
func WithSkipPollingDelay(ctx context.Context) context.Context {
	return context.WithValue(ctx, skipPollingDelayKey, true)
}

// ShouldSkipPollingDelay returns true if the context has the skip polling delay flag set.
func ShouldSkipPollingDelay(ctx context.Context) bool {
	if v, ok := ctx.Value(skipPollingDelayKey).(bool); ok {
		return v
	}
	return false
}

// WithInitialPollingDelay returns a new context with the initial polling delay value set.
// This is used to signal to PollUntilDone the initial delay before the first polling attempt.
// It's intended to ignore if value is less than or equal to 0
func WithInitialPollingDelay(ctx context.Context, value time.Duration) context.Context {
	if value <= 0 {
		return ctx
	}
	return context.WithValue(ctx, initialPollingDelayKey, value)
}

// GetInitialPollingDelay returns the initial polling delay value if set.
func GetInitialPollingDelay(ctx context.Context) time.Duration {
	if v, ok := ctx.Value(initialPollingDelayKey).(time.Duration); ok {
		return v
	}
	return 0
}
