package daprcomponentresiliencypolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprComponentResiliencyPolicyConfiguration struct {
	CircuitBreakerPolicy *DaprComponentResiliencyPolicyCircuitBreakerPolicyConfiguration `json:"circuitBreakerPolicy,omitempty"`
	HTTPRetryPolicy      *DaprComponentResiliencyPolicyHTTPRetryPolicyConfiguration      `json:"httpRetryPolicy,omitempty"`
	TimeoutPolicy        *DaprComponentResiliencyPolicyTimeoutPolicyConfiguration        `json:"timeoutPolicy,omitempty"`
}
