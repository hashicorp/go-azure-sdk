package hcrpreports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReportProperties struct {
	ConfigurationProfile *string           `json:"configurationProfile,omitempty"`
	Duration             *string           `json:"duration,omitempty"`
	EndTime              *string           `json:"endTime,omitempty"`
	Error                *ErrorDetail      `json:"error,omitempty"`
	LastModifiedTime     *string           `json:"lastModifiedTime,omitempty"`
	ReportFormatVersion  *string           `json:"reportFormatVersion,omitempty"`
	Resources            *[]ReportResource `json:"resources,omitempty"`
	StartTime            *string           `json:"startTime,omitempty"`
	Status               *string           `json:"status,omitempty"`
	Type                 *string           `json:"type,omitempty"`
}
