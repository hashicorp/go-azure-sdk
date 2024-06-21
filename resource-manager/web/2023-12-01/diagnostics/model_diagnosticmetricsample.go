package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticMetricSample struct {
	IsAggregated *bool    `json:"isAggregated,omitempty"`
	Maximum      *float64 `json:"maximum,omitempty"`
	Minimum      *float64 `json:"minimum,omitempty"`
	RoleInstance *string  `json:"roleInstance,omitempty"`
	Timestamp    *string  `json:"timestamp,omitempty"`
	Total        *float64 `json:"total,omitempty"`
}
