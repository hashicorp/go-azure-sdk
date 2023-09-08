package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonTypeOperationPredicate struct {
	Class     *string
	ODataType *string
	Subclass  *string
}

func (p PersonTypeOperationPredicate) Matches(input PersonType) bool {

	if p.Class != nil && (input.Class == nil || *p.Class != *input.Class) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Subclass != nil && (input.Subclass == nil || *p.Subclass != *input.Subclass) {
		return false
	}

	return true
}
