package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticMetricSet struct {
	EndTime   *string                   `json:"endTime,omitempty"`
	Name      *string                   `json:"name,omitempty"`
	StartTime *string                   `json:"startTime,omitempty"`
	TimeGrain *string                   `json:"timeGrain,omitempty"`
	Unit      *string                   `json:"unit,omitempty"`
	Values    *[]DiagnosticMetricSample `json:"values,omitempty"`
}
