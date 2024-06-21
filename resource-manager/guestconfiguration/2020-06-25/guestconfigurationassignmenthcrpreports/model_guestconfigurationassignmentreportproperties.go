package guestconfigurationassignmenthcrpreports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationAssignmentReportProperties struct {
	Assignment       *AssignmentInfo          `json:"assignment,omitempty"`
	ComplianceStatus *ComplianceStatus        `json:"complianceStatus,omitempty"`
	Details          *AssignmentReportDetails `json:"details,omitempty"`
	EndTime          *string                  `json:"endTime,omitempty"`
	ReportId         *string                  `json:"reportId,omitempty"`
	StartTime        *string                  `json:"startTime,omitempty"`
	VM               *VMInfo                  `json:"vm,omitempty"`
	VMSSResourceId   *string                  `json:"vmssResourceId,omitempty"`
}
