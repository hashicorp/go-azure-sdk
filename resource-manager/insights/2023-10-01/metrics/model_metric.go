package metrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Metric struct {
	DisplayDescription *string             `json:"displayDescription,omitempty"`
	ErrorCode          *string             `json:"errorCode,omitempty"`
	ErrorMessage       *string             `json:"errorMessage,omitempty"`
	Id                 string              `json:"id"`
	Name               LocalizableString   `json:"name"`
	Timeseries         []TimeSeriesElement `json:"timeseries"`
	Type               string              `json:"type"`
	Unit               MetricUnit          `json:"unit"`
}
