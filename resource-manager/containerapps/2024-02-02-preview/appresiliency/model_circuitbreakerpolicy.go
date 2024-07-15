package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CircuitBreakerPolicy struct {
	ConsecutiveErrors  *int64 `json:"consecutiveErrors,omitempty"`
	IntervalInSeconds  *int64 `json:"intervalInSeconds,omitempty"`
	MaxEjectionPercent *int64 `json:"maxEjectionPercent,omitempty"`
}
