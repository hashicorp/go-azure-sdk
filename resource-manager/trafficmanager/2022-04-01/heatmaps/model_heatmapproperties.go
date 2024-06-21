package heatmaps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HeatMapProperties struct {
	EndTime      *string            `json:"endTime,omitempty"`
	Endpoints    *[]HeatMapEndpoint `json:"endpoints,omitempty"`
	StartTime    *string            `json:"startTime,omitempty"`
	TrafficFlows *[]TrafficFlow     `json:"trafficFlows,omitempty"`
}
