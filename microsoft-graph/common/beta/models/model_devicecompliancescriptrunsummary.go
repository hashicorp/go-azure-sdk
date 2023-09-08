package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRunSummary struct {
	DetectionScriptErrorDeviceCount   *int64  `json:"detectionScriptErrorDeviceCount,omitempty"`
	DetectionScriptPendingDeviceCount *int64  `json:"detectionScriptPendingDeviceCount,omitempty"`
	Id                                *string `json:"id,omitempty"`
	IssueDetectedDeviceCount          *int64  `json:"issueDetectedDeviceCount,omitempty"`
	LastScriptRunDateTime             *string `json:"lastScriptRunDateTime,omitempty"`
	NoIssueDetectedDeviceCount        *int64  `json:"noIssueDetectedDeviceCount,omitempty"`
	ODataType                         *string `json:"@odata.type,omitempty"`
}
