package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostalAddressTypeOperationPredicate struct {
	City              *string
	CountryLetterCode *string
	ODataType         *string
	PostalCode        *string
	State             *string
	Street            *string
}

func (p PostalAddressTypeOperationPredicate) Matches(input PostalAddressType) bool {

	if p.City != nil && (input.City == nil || *p.City != *input.City) {
		return false
	}

	if p.CountryLetterCode != nil && (input.CountryLetterCode == nil || *p.CountryLetterCode != *input.CountryLetterCode) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PostalCode != nil && (input.PostalCode == nil || *p.PostalCode != *input.PostalCode) {
		return false
	}

	if p.State != nil && (input.State == nil || *p.State != *input.State) {
		return false
	}

	if p.Street != nil && (input.Street == nil || *p.Street != *input.Street) {
		return false
	}

	return true
}
