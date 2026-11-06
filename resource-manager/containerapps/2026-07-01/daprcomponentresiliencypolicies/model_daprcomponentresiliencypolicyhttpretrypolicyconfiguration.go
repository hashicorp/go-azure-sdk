package daprcomponentresiliencypolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprComponentResiliencyPolicyHTTPRetryPolicyConfiguration struct {
	MaxRetries   *int64                                                      `json:"maxRetries,omitempty"`
	RetryBackOff *DaprComponentResiliencyPolicyHTTPRetryBackOffConfiguration `json:"retryBackOff,omitempty"`
}
