package sqlpoolsworkloadgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadGroupProperties struct {
	Importance                   *string  `json:"importance,omitempty"`
	MaxResourcePercent           int64    `json:"maxResourcePercent"`
	MaxResourcePercentPerRequest *float64 `json:"maxResourcePercentPerRequest,omitempty"`
	MinResourcePercent           int64    `json:"minResourcePercent"`
	MinResourcePercentPerRequest float64  `json:"minResourcePercentPerRequest"`
	QueryExecutionTimeout        *int64   `json:"queryExecutionTimeout,omitempty"`
}
