package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigManagerPolicySummaryOperationPredicate struct {
	CompliantDeviceCount    *int64
	EnforcedDeviceCount     *int64
	FailedDeviceCount       *int64
	NonCompliantDeviceCount *int64
	ODataType               *string
	PendingDeviceCount      *int64
	TargetedDeviceCount     *int64
}

func (p ConfigManagerPolicySummaryOperationPredicate) Matches(input ConfigManagerPolicySummary) bool {

	if p.CompliantDeviceCount != nil && (input.CompliantDeviceCount == nil || *p.CompliantDeviceCount != *input.CompliantDeviceCount) {
		return false
	}

	if p.EnforcedDeviceCount != nil && (input.EnforcedDeviceCount == nil || *p.EnforcedDeviceCount != *input.EnforcedDeviceCount) {
		return false
	}

	if p.FailedDeviceCount != nil && (input.FailedDeviceCount == nil || *p.FailedDeviceCount != *input.FailedDeviceCount) {
		return false
	}

	if p.NonCompliantDeviceCount != nil && (input.NonCompliantDeviceCount == nil || *p.NonCompliantDeviceCount != *input.NonCompliantDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PendingDeviceCount != nil && (input.PendingDeviceCount == nil || *p.PendingDeviceCount != *input.PendingDeviceCount) {
		return false
	}

	if p.TargetedDeviceCount != nil && (input.TargetedDeviceCount == nil || *p.TargetedDeviceCount != *input.TargetedDeviceCount) {
		return false
	}

	return true
}
