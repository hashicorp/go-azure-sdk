package datacollectionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogsQuotaSpec struct {
	MaxRequestsPerMinute *string `json:"maxRequestsPerMinute,omitempty"`
	MaxSizePerMinuteInGB *string `json:"maxSizePerMinuteInGB,omitempty"`
}
