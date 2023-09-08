package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RankedEmailAddressOperationPredicate struct {
	Address   *string
	ODataType *string
}

func (p RankedEmailAddressOperationPredicate) Matches(input RankedEmailAddress) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
