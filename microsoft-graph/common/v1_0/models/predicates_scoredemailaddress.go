package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScoredEmailAddressOperationPredicate struct {
	Address   *string
	ItemId    *string
	ODataType *string
}

func (p ScoredEmailAddressOperationPredicate) Matches(input ScoredEmailAddress) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.ItemId != nil && (input.ItemId == nil || *p.ItemId != *input.ItemId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
