package staticsitesasyncoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSitesOperationStatusProperties struct {
	EndTime              *string        `json:"endTime,omitempty"`
	Error                *ErrorResponse `json:"error,omitempty"`
	StartTime            *string        `json:"startTime,omitempty"`
	StaticSiteProperties *StaticSite    `json:"staticSiteProperties,omitempty"`
	Status               *string        `json:"status,omitempty"`
}
