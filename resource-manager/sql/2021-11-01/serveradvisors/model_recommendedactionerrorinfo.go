package serveradvisors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionErrorInfo struct {
	ErrorCode   *string      `json:"errorCode,omitempty"`
	IsRetryable *IsRetryable `json:"isRetryable,omitempty"`
}
