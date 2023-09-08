package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestOperationPredicate struct {
	CompletedDate      *string
	CreatedDateTime    *string
	ExpirationDateTime *string
	Id                 *string
	IsValidationOnly   *bool
	Justification      *string
	ODataType          *string
	RequestState       *string
	RequestStatus      *string
	RequestType        *string
}

func (p AccessPackageAssignmentRequestOperationPredicate) Matches(input AccessPackageAssignmentRequest) bool {

	if p.CompletedDate != nil && (input.CompletedDate == nil || *p.CompletedDate != *input.CompletedDate) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsValidationOnly != nil && (input.IsValidationOnly == nil || *p.IsValidationOnly != *input.IsValidationOnly) {
		return false
	}

	if p.Justification != nil && (input.Justification == nil || *p.Justification != *input.Justification) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestState != nil && (input.RequestState == nil || *p.RequestState != *input.RequestState) {
		return false
	}

	if p.RequestStatus != nil && (input.RequestStatus == nil || *p.RequestStatus != *input.RequestStatus) {
		return false
	}

	if p.RequestType != nil && (input.RequestType == nil || *p.RequestType != *input.RequestType) {
		return false
	}

	return true
}
