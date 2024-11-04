package onlinedeployment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataCollector struct {
	Collections    map[string]Collection `json:"collections"`
	RequestLogging *RequestLogging       `json:"requestLogging,omitempty"`
	RollingRate    *RollingRateType      `json:"rollingRate,omitempty"`
}
