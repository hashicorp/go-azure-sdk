package daprcomponentresiliencypolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprComponentResiliencyPolicyCircuitBreakerPolicyConfiguration struct {
	ConsecutiveErrors *int64 `json:"consecutiveErrors,omitempty"`
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty"`
	TimeoutInSeconds  *int64 `json:"timeoutInSeconds,omitempty"`
}
