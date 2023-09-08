package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyOperationPredicate struct {
	BusinessProfileId *string
	DisplayName       *string
	Id                *string
	Name              *string
	ODataType         *string
	SystemVersion     *string
}

func (p CompanyOperationPredicate) Matches(input Company) bool {

	if p.BusinessProfileId != nil && (input.BusinessProfileId == nil || *p.BusinessProfileId != *input.BusinessProfileId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SystemVersion != nil && (input.SystemVersion == nil || *p.SystemVersion != *input.SystemVersion) {
		return false
	}

	return true
}
