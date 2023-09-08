package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BooleanColumnOperationPredicate struct {
	ODataType *string
}

func (p BooleanColumnOperationPredicate) Matches(input BooleanColumn) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
