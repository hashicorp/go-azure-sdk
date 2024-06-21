package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AbnormalTimePeriod struct {
	EndTime   *string                       `json:"endTime,omitempty"`
	Events    *[]DetectorAbnormalTimePeriod `json:"events,omitempty"`
	Solutions *[]Solution                   `json:"solutions,omitempty"`
	StartTime *string                       `json:"startTime,omitempty"`
}
