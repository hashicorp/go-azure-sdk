package guestconfigurationconnectedvmwarevsphereassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReport struct {
	Assignment       *AssignmentInfo             `json:"assignment,omitempty"`
	ComplianceStatus *ComplianceStatus           `json:"complianceStatus,omitempty"`
	EndTime          *string                     `json:"endTime,omitempty"`
	Id               *string                     `json:"id,omitempty"`
	OperationType    *Type                       `json:"operationType,omitempty"`
	ReportId         *string                     `json:"reportId,omitempty"`
	Resources        *[]AssignmentReportResource `json:"resources,omitempty"`
	StartTime        *string                     `json:"startTime,omitempty"`
	VM               *VMInfo                     `json:"vm,omitempty"`
}
