package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeoutPolicy struct {
	ConnectionTimeoutInSeconds *int64 `json:"connectionTimeoutInSeconds,omitempty"`
	ResponseTimeoutInSeconds   *int64 `json:"responseTimeoutInSeconds,omitempty"`
}
