package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedDataLocationOperationPredicate struct {
	AppId     *string
	Domain    *string
	Id        *string
	IsDefault *bool
	Location  *string
	ODataType *string
}

func (p AllowedDataLocationOperationPredicate) Matches(input AllowedDataLocation) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.Domain != nil && (input.Domain == nil || *p.Domain != *input.Domain) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefault != nil && (input.IsDefault == nil || *p.IsDefault != *input.IsDefault) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
