package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataReferenceValueOperationPredicate struct {
	Code      *string
	ODataType *string
}

func (p IndustryDataReferenceValueOperationPredicate) Matches(input IndustryDataReferenceValue) bool {

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
