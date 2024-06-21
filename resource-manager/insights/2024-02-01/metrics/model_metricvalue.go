package metrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricValue struct {
	Average   *float64 `json:"average,omitempty"`
	Count     *float64 `json:"count,omitempty"`
	Maximum   *float64 `json:"maximum,omitempty"`
	Minimum   *float64 `json:"minimum,omitempty"`
	TimeStamp string   `json:"timeStamp"`
	Total     *float64 `json:"total,omitempty"`
}
