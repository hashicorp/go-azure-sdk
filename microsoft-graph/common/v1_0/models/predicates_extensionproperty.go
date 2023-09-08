package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionPropertyOperationPredicate struct {
	AppDisplayName         *string
	DataType               *string
	DeletedDateTime        *string
	Id                     *string
	IsSyncedFromOnPremises *bool
	Name                   *string
	ODataType              *string
}

func (p ExtensionPropertyOperationPredicate) Matches(input ExtensionProperty) bool {

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.DataType != nil && (input.DataType == nil || *p.DataType != *input.DataType) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSyncedFromOnPremises != nil && (input.IsSyncedFromOnPremises == nil || *p.IsSyncedFromOnPremises != *input.IsSyncedFromOnPremises) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
