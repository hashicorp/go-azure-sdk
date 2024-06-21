package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerfMonSet struct {
	EndTime   *string          `json:"endTime,omitempty"`
	Name      *string          `json:"name,omitempty"`
	StartTime *string          `json:"startTime,omitempty"`
	TimeGrain *string          `json:"timeGrain,omitempty"`
	Values    *[]PerfMonSample `json:"values,omitempty"`
}
