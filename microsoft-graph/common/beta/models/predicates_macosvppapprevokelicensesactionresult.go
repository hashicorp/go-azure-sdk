package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppRevokeLicensesActionResultOperationPredicate struct {
	ActionName          *string
	FailedLicensesCount *int64
	LastUpdatedDateTime *string
	ManagedDeviceId     *string
	ODataType           *string
	StartDateTime       *string
	TotalLicensesCount  *int64
	UserId              *string
}

func (p MacOsVppAppRevokeLicensesActionResultOperationPredicate) Matches(input MacOsVppAppRevokeLicensesActionResult) bool {

	if p.ActionName != nil && (input.ActionName == nil || *p.ActionName != *input.ActionName) {
		return false
	}

	if p.FailedLicensesCount != nil && (input.FailedLicensesCount == nil || *p.FailedLicensesCount != *input.FailedLicensesCount) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.TotalLicensesCount != nil && (input.TotalLicensesCount == nil || *p.TotalLicensesCount != *input.TotalLicensesCount) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
