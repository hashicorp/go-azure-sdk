package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceComplianceLocalActionLockDeviceOperationPredicate struct {
	GracePeriodInMinutes *int64
	Id                   *string
	ODataType            *string
}

func (p AndroidDeviceComplianceLocalActionLockDeviceOperationPredicate) Matches(input AndroidDeviceComplianceLocalActionLockDevice) bool {

	if p.GracePeriodInMinutes != nil && (input.GracePeriodInMinutes == nil || *p.GracePeriodInMinutes != *input.GracePeriodInMinutes) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
