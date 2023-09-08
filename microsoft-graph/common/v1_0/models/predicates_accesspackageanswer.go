package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAnswerOperationPredicate struct {
	DisplayValue *string
	ODataType    *string
}

func (p AccessPackageAnswerOperationPredicate) Matches(input AccessPackageAnswer) bool {

	if p.DisplayValue != nil && (input.DisplayValue == nil || *p.DisplayValue != *input.DisplayValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
