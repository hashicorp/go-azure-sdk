package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentDeviceStateSummaryOperationPredicate struct {
	ConflictCount              *int64
	ErrorCount                 *int64
	FailedCount                *int64
	Id                         *string
	NotApplicableCount         *int64
	NotApplicablePlatformCount *int64
	ODataType                  *string
	SuccessCount               *int64
}

func (p DeviceManagementIntentDeviceStateSummaryOperationPredicate) Matches(input DeviceManagementIntentDeviceStateSummary) bool {

	if p.ConflictCount != nil && (input.ConflictCount == nil || *p.ConflictCount != *input.ConflictCount) {
		return false
	}

	if p.ErrorCount != nil && (input.ErrorCount == nil || *p.ErrorCount != *input.ErrorCount) {
		return false
	}

	if p.FailedCount != nil && (input.FailedCount == nil || *p.FailedCount != *input.FailedCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.NotApplicableCount != nil && (input.NotApplicableCount == nil || *p.NotApplicableCount != *input.NotApplicableCount) {
		return false
	}

	if p.NotApplicablePlatformCount != nil && (input.NotApplicablePlatformCount == nil || *p.NotApplicablePlatformCount != *input.NotApplicablePlatformCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessCount != nil && (input.SuccessCount == nil || *p.SuccessCount != *input.SuccessCount) {
		return false
	}

	return true
}
