package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPv4RangeOperationPredicate struct {
	LowerAddress *string
	ODataType    *string
	UpperAddress *string
}

func (p IPv4RangeOperationPredicate) Matches(input IPv4Range) bool {

	if p.LowerAddress != nil && (input.LowerAddress == nil || *p.LowerAddress != *input.LowerAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UpperAddress != nil && (input.UpperAddress == nil || *p.UpperAddress != *input.UpperAddress) {
		return false
	}

	return true
}
