package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRequestOperationPredicate struct {
	CatalogId          *string
	ExecuteImmediately *bool
	ExpirationDateTime *string
	Id                 *string
	IsValidationOnly   *bool
	Justification      *string
	ODataType          *string
	RequestState       *string
	RequestStatus      *string
	RequestType        *string
}

func (p AccessPackageResourceRequestOperationPredicate) Matches(input AccessPackageResourceRequest) bool {

	if p.CatalogId != nil && (input.CatalogId == nil || *p.CatalogId != *input.CatalogId) {
		return false
	}

	if p.ExecuteImmediately != nil && (input.ExecuteImmediately == nil || *p.ExecuteImmediately != *input.ExecuteImmediately) {
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
