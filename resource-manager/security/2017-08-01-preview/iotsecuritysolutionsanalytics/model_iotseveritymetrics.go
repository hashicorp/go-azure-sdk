package iotsecuritysolutionsanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSeverityMetrics struct {
	High   *int64 `json:"high,omitempty"`
	Low    *int64 `json:"low,omitempty"`
	Medium *int64 `json:"medium,omitempty"`
}
