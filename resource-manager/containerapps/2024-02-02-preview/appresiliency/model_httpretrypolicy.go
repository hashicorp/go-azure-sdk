package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HTTPRetryPolicy struct {
	Matches      *HTTPRetryPolicyMatches      `json:"matches,omitempty"`
	MaxRetries   *int64                       `json:"maxRetries,omitempty"`
	RetryBackOff *HTTPRetryPolicyRetryBackOff `json:"retryBackOff,omitempty"`
}
