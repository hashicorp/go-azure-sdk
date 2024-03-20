package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365Source struct {
	AllowedGroups            *interface{} `json:"allowedGroups,omitempty"`
	DateFilterColumn         *interface{} `json:"dateFilterColumn,omitempty"`
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	EndTime                  *interface{} `json:"endTime,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	OutputColumns            *interface{} `json:"outputColumns,omitempty"`
	SourceRetryCount         *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
	StartTime                *interface{} `json:"startTime,omitempty"`
	Type                     string       `json:"type"`
	UserScopeFilterUri       *interface{} `json:"userScopeFilterUri,omitempty"`
}
