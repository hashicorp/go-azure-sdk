package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingStateSummaryOperationPredicate struct {
	CompliantDeviceCount     *int64
	ConflictDeviceCount      *int64
	ErrorDeviceCount         *int64
	Id                       *string
	NonCompliantDeviceCount  *int64
	NotApplicableDeviceCount *int64
	NotAssignedDeviceCount   *int64
	ODataType                *string
	RemediatedDeviceCount    *int64
	UnknownDeviceCount       *int64
}

func (p AdvancedThreatProtectionOnboardingStateSummaryOperationPredicate) Matches(input AdvancedThreatProtectionOnboardingStateSummary) bool {

	if p.CompliantDeviceCount != nil && (input.CompliantDeviceCount == nil || *p.CompliantDeviceCount != *input.CompliantDeviceCount) {
		return false
	}

	if p.ConflictDeviceCount != nil && (input.ConflictDeviceCount == nil || *p.ConflictDeviceCount != *input.ConflictDeviceCount) {
		return false
	}

	if p.ErrorDeviceCount != nil && (input.ErrorDeviceCount == nil || *p.ErrorDeviceCount != *input.ErrorDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.NonCompliantDeviceCount != nil && (input.NonCompliantDeviceCount == nil || *p.NonCompliantDeviceCount != *input.NonCompliantDeviceCount) {
		return false
	}

	if p.NotApplicableDeviceCount != nil && (input.NotApplicableDeviceCount == nil || *p.NotApplicableDeviceCount != *input.NotApplicableDeviceCount) {
		return false
	}

	if p.NotAssignedDeviceCount != nil && (input.NotAssignedDeviceCount == nil || *p.NotAssignedDeviceCount != *input.NotAssignedDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediatedDeviceCount != nil && (input.RemediatedDeviceCount == nil || *p.RemediatedDeviceCount != *input.RemediatedDeviceCount) {
		return false
	}

	if p.UnknownDeviceCount != nil && (input.UnknownDeviceCount == nil || *p.UnknownDeviceCount != *input.UnknownDeviceCount) {
		return false
	}

	return true
}
