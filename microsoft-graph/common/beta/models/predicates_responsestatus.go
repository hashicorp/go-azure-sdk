package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponseStatusOperationPredicate struct {
	ODataType *string
	Time      *string
}

func (p ResponseStatusOperationPredicate) Matches(input ResponseStatus) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Time != nil && (input.Time == nil || *p.Time != *input.Time) {
		return false
	}

	return true
}
