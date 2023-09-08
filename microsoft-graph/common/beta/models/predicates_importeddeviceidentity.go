package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityOperationPredicate struct {
	CreatedDateTime          *string
	Description              *string
	Id                       *string
	ImportedDeviceIdentifier *string
	LastContactedDateTime    *string
	LastModifiedDateTime     *string
	ODataType                *string
}

func (p ImportedDeviceIdentityOperationPredicate) Matches(input ImportedDeviceIdentity) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ImportedDeviceIdentifier != nil && (input.ImportedDeviceIdentifier == nil || *p.ImportedDeviceIdentifier != *input.ImportedDeviceIdentifier) {
		return false
	}

	if p.LastContactedDateTime != nil && (input.LastContactedDateTime == nil || *p.LastContactedDateTime != *input.LastContactedDateTime) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
