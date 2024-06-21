package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorAbnormalTimePeriod struct {
	EndTime   *string            `json:"endTime,omitempty"`
	Message   *string            `json:"message,omitempty"`
	MetaData  *[][]NameValuePair `json:"metaData,omitempty"`
	Priority  *float64           `json:"priority,omitempty"`
	Solutions *[]Solution        `json:"solutions,omitempty"`
	Source    *string            `json:"source,omitempty"`
	StartTime *string            `json:"startTime,omitempty"`
	Type      *IssueType         `json:"type,omitempty"`
}
