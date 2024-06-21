package guestconfigurationassignmenthcrpreports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReportDetails struct {
	ComplianceStatus *ComplianceStatus           `json:"complianceStatus,omitempty"`
	EndTime          *string                     `json:"endTime,omitempty"`
	JobId            *string                     `json:"jobId,omitempty"`
	OperationType    *Type                       `json:"operationType,omitempty"`
	Resources        *[]AssignmentReportResource `json:"resources,omitempty"`
	StartTime        *string                     `json:"startTime,omitempty"`
}
