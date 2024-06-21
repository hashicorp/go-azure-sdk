package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityGetInsightsParameters struct {
	AddDefaultExtendedTimeRange *bool     `json:"addDefaultExtendedTimeRange,omitempty"`
	EndTime                     string    `json:"endTime"`
	InsightQueryIds             *[]string `json:"insightQueryIds,omitempty"`
	StartTime                   string    `json:"startTime"`
}
