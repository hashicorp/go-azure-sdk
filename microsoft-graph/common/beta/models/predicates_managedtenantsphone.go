package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsPhoneOperationPredicate struct {
	ODataType   *string
	PhoneNumber *string
}

func (p ManagedTenantsPhoneOperationPredicate) Matches(input ManagedTenantsPhone) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PhoneNumber != nil && (input.PhoneNumber == nil || *p.PhoneNumber != *input.PhoneNumber) {
		return false
	}

	return true
}
