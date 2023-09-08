package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentPolicyOperationPredicate struct {
	AccessPackageId    *string
	CanExtend          *bool
	CreatedBy          *string
	CreatedDateTime    *string
	Description        *string
	DisplayName        *string
	DurationInDays     *int64
	ExpirationDateTime *string
	Id                 *string
	ModifiedBy         *string
	ModifiedDateTime   *string
	ODataType          *string
}

func (p AccessPackageAssignmentPolicyOperationPredicate) Matches(input AccessPackageAssignmentPolicy) bool {

	if p.AccessPackageId != nil && (input.AccessPackageId == nil || *p.AccessPackageId != *input.AccessPackageId) {
		return false
	}

	if p.CanExtend != nil && (input.CanExtend == nil || *p.CanExtend != *input.CanExtend) {
		return false
	}

	if p.CreatedBy != nil && (input.CreatedBy == nil || *p.CreatedBy != *input.CreatedBy) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DurationInDays != nil && (input.DurationInDays == nil || *p.DurationInDays != *input.DurationInDays) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ModifiedBy != nil && (input.ModifiedBy == nil || *p.ModifiedBy != *input.ModifiedBy) {
		return false
	}

	if p.ModifiedDateTime != nil && (input.ModifiedDateTime == nil || *p.ModifiedDateTime != *input.ModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
