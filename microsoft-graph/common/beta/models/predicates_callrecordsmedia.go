package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaOperationPredicate struct {
	Label     *string
	ODataType *string
}

func (p CallRecordsMediaOperationPredicate) Matches(input CallRecordsMedia) bool {

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
