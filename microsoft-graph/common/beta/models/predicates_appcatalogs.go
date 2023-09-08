package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCatalogsOperationPredicate struct {
	ODataType *string
}

func (p AppCatalogsOperationPredicate) Matches(input AppCatalogs) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
