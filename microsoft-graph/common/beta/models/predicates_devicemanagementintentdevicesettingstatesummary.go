package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate struct {
	CompliantCount     *int64
	ConflictCount      *int64
	ErrorCount         *int64
	Id                 *string
	NonCompliantCount  *int64
	NotApplicableCount *int64
	ODataType          *string
	RemediatedCount    *int64
	SettingName        *string
}

func (p DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate) Matches(input DeviceManagementIntentDeviceSettingStateSummary) bool {

	if p.CompliantCount != nil && (input.CompliantCount == nil || *p.CompliantCount != *input.CompliantCount) {
		return false
	}

	if p.ConflictCount != nil && (input.ConflictCount == nil || *p.ConflictCount != *input.ConflictCount) {
		return false
	}

	if p.ErrorCount != nil && (input.ErrorCount == nil || *p.ErrorCount != *input.ErrorCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.NonCompliantCount != nil && (input.NonCompliantCount == nil || *p.NonCompliantCount != *input.NonCompliantCount) {
		return false
	}

	if p.NotApplicableCount != nil && (input.NotApplicableCount == nil || *p.NotApplicableCount != *input.NotApplicableCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediatedCount != nil && (input.RemediatedCount == nil || *p.RemediatedCount != *input.RemediatedCount) {
		return false
	}

	if p.SettingName != nil && (input.SettingName == nil || *p.SettingName != *input.SettingName) {
		return false
	}

	return true
}
