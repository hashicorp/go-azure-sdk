package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyDetailOperationPredicate struct {
	Department     *string
	DisplayName    *string
	ODataType      *string
	OfficeLocation *string
	Pronunciation  *string
	WebUrl         *string
}

func (p CompanyDetailOperationPredicate) Matches(input CompanyDetail) bool {

	if p.Department != nil && (input.Department == nil || *p.Department != *input.Department) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OfficeLocation != nil && (input.OfficeLocation == nil || *p.OfficeLocation != *input.OfficeLocation) {
		return false
	}

	if p.Pronunciation != nil && (input.Pronunciation == nil || *p.Pronunciation != *input.Pronunciation) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
