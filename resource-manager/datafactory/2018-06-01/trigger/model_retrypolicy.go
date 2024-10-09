package trigger

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetryPolicy struct {
	Count             *int64 `json:"count,omitempty"`
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty"`
}
