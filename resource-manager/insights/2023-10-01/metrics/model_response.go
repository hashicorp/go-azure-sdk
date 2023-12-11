package metrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Response struct {
	Cost           *float64 `json:"cost,omitempty"`
	Interval       *string  `json:"interval,omitempty"`
	Namespace      *string  `json:"namespace,omitempty"`
	Resourceregion *string  `json:"resourceregion,omitempty"`
	Timespan       string   `json:"timespan"`
	Value          []Metric `json:"value"`
}
