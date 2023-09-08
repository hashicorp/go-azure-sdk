package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2KeyRestrictionsOperationPredicate struct {
	IsEnforced *bool
	ODataType  *string
}

func (p Fido2KeyRestrictionsOperationPredicate) Matches(input Fido2KeyRestrictions) bool {

	if p.IsEnforced != nil && (input.IsEnforced == nil || *p.IsEnforced != *input.IsEnforced) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
