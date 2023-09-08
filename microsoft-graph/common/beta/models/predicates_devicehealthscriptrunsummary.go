package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRunSummaryOperationPredicate struct {
	DetectionScriptErrorDeviceCount         *int64
	DetectionScriptNotApplicableDeviceCount *int64
	DetectionScriptPendingDeviceCount       *int64
	Id                                      *string
	IssueDetectedDeviceCount                *int64
	IssueRemediatedCumulativeDeviceCount    *int64
	IssueRemediatedDeviceCount              *int64
	IssueReoccurredDeviceCount              *int64
	LastScriptRunDateTime                   *string
	NoIssueDetectedDeviceCount              *int64
	ODataType                               *string
	RemediationScriptErrorDeviceCount       *int64
	RemediationSkippedDeviceCount           *int64
}

func (p DeviceHealthScriptRunSummaryOperationPredicate) Matches(input DeviceHealthScriptRunSummary) bool {

	if p.DetectionScriptErrorDeviceCount != nil && (input.DetectionScriptErrorDeviceCount == nil || *p.DetectionScriptErrorDeviceCount != *input.DetectionScriptErrorDeviceCount) {
		return false
	}

	if p.DetectionScriptNotApplicableDeviceCount != nil && (input.DetectionScriptNotApplicableDeviceCount == nil || *p.DetectionScriptNotApplicableDeviceCount != *input.DetectionScriptNotApplicableDeviceCount) {
		return false
	}

	if p.DetectionScriptPendingDeviceCount != nil && (input.DetectionScriptPendingDeviceCount == nil || *p.DetectionScriptPendingDeviceCount != *input.DetectionScriptPendingDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IssueDetectedDeviceCount != nil && (input.IssueDetectedDeviceCount == nil || *p.IssueDetectedDeviceCount != *input.IssueDetectedDeviceCount) {
		return false
	}

	if p.IssueRemediatedCumulativeDeviceCount != nil && (input.IssueRemediatedCumulativeDeviceCount == nil || *p.IssueRemediatedCumulativeDeviceCount != *input.IssueRemediatedCumulativeDeviceCount) {
		return false
	}

	if p.IssueRemediatedDeviceCount != nil && (input.IssueRemediatedDeviceCount == nil || *p.IssueRemediatedDeviceCount != *input.IssueRemediatedDeviceCount) {
		return false
	}

	if p.IssueReoccurredDeviceCount != nil && (input.IssueReoccurredDeviceCount == nil || *p.IssueReoccurredDeviceCount != *input.IssueReoccurredDeviceCount) {
		return false
	}

	if p.LastScriptRunDateTime != nil && (input.LastScriptRunDateTime == nil || *p.LastScriptRunDateTime != *input.LastScriptRunDateTime) {
		return false
	}

	if p.NoIssueDetectedDeviceCount != nil && (input.NoIssueDetectedDeviceCount == nil || *p.NoIssueDetectedDeviceCount != *input.NoIssueDetectedDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediationScriptErrorDeviceCount != nil && (input.RemediationScriptErrorDeviceCount == nil || *p.RemediationScriptErrorDeviceCount != *input.RemediationScriptErrorDeviceCount) {
		return false
	}

	if p.RemediationSkippedDeviceCount != nil && (input.RemediationSkippedDeviceCount == nil || *p.RemediationSkippedDeviceCount != *input.RemediationSkippedDeviceCount) {
		return false
	}

	return true
}
