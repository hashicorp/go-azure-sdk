package guestconfigurationassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMSSVMInfo struct {
	ComplianceStatus      *ComplianceStatus `json:"complianceStatus,omitempty"`
	LastComplianceChecked *string           `json:"lastComplianceChecked,omitempty"`
	LatestReportId        *string           `json:"latestReportId,omitempty"`
	VMId                  *string           `json:"vmId,omitempty"`
	VMResourceId          *string           `json:"vmResourceId,omitempty"`
}
