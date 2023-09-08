package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityOperationPredicate struct {
	CreatedDateTime                              *string
	Description                                  *string
	Id                                           *string
	IsDeleted                                    *bool
	IsSupervised                                 *bool
	LastContactedDateTime                        *string
	ODataType                                    *string
	RequestedEnrollmentProfileAssignmentDateTime *string
	RequestedEnrollmentProfileId                 *string
	SerialNumber                                 *string
}

func (p ImportedAppleDeviceIdentityOperationPredicate) Matches(input ImportedAppleDeviceIdentity) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDeleted != nil && (input.IsDeleted == nil || *p.IsDeleted != *input.IsDeleted) {
		return false
	}

	if p.IsSupervised != nil && (input.IsSupervised == nil || *p.IsSupervised != *input.IsSupervised) {
		return false
	}

	if p.LastContactedDateTime != nil && (input.LastContactedDateTime == nil || *p.LastContactedDateTime != *input.LastContactedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestedEnrollmentProfileAssignmentDateTime != nil && (input.RequestedEnrollmentProfileAssignmentDateTime == nil || *p.RequestedEnrollmentProfileAssignmentDateTime != *input.RequestedEnrollmentProfileAssignmentDateTime) {
		return false
	}

	if p.RequestedEnrollmentProfileId != nil && (input.RequestedEnrollmentProfileId == nil || *p.RequestedEnrollmentProfileId != *input.RequestedEnrollmentProfileId) {
		return false
	}

	if p.SerialNumber != nil && (input.SerialNumber == nil || *p.SerialNumber != *input.SerialNumber) {
		return false
	}

	return true
}
