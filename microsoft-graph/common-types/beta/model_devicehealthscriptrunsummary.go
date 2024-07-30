package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRunSummary struct {
	DetectionScriptErrorDeviceCount         *int64  `json:"detectionScriptErrorDeviceCount,omitempty"`
	DetectionScriptNotApplicableDeviceCount *int64  `json:"detectionScriptNotApplicableDeviceCount,omitempty"`
	DetectionScriptPendingDeviceCount       *int64  `json:"detectionScriptPendingDeviceCount,omitempty"`
	Id                                      *string `json:"id,omitempty"`
	IssueDetectedDeviceCount                *int64  `json:"issueDetectedDeviceCount,omitempty"`
	IssueRemediatedCumulativeDeviceCount    *int64  `json:"issueRemediatedCumulativeDeviceCount,omitempty"`
	IssueRemediatedDeviceCount              *int64  `json:"issueRemediatedDeviceCount,omitempty"`
	IssueReoccurredDeviceCount              *int64  `json:"issueReoccurredDeviceCount,omitempty"`
	LastScriptRunDateTime                   *string `json:"lastScriptRunDateTime,omitempty"`
	NoIssueDetectedDeviceCount              *int64  `json:"noIssueDetectedDeviceCount,omitempty"`
	ODataType                               *string `json:"@odata.type,omitempty"`
	RemediationScriptErrorDeviceCount       *int64  `json:"remediationScriptErrorDeviceCount,omitempty"`
	RemediationSkippedDeviceCount           *int64  `json:"remediationSkippedDeviceCount,omitempty"`
}
