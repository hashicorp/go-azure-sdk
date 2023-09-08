package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityResourceOperationPredicate struct {
	ODataType *string
	Resource  *string
}

func (p SecurityResourceOperationPredicate) Matches(input SecurityResource) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Resource != nil && (input.Resource == nil || *p.Resource != *input.Resource) {
		return false
	}

	return true
}
