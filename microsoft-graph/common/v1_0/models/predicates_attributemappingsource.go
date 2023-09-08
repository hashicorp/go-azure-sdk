package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingSourceOperationPredicate struct {
	Expression *string
	Name       *string
	ODataType  *string
}

func (p AttributeMappingSourceOperationPredicate) Matches(input AttributeMappingSource) bool {

	if p.Expression != nil && (input.Expression == nil || *p.Expression != *input.Expression) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
