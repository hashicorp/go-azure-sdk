package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BundleOperationPredicate struct {
	ChildCount *int64
	ODataType  *string
}

func (p BundleOperationPredicate) Matches(input Bundle) bool {

	if p.ChildCount != nil && (input.ChildCount == nil || *p.ChildCount != *input.ChildCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
