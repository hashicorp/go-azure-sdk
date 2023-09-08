package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalculatedColumnOperationPredicate struct {
	Format     *string
	Formula    *string
	ODataType  *string
	OutputType *string
}

func (p CalculatedColumnOperationPredicate) Matches(input CalculatedColumn) bool {

	if p.Format != nil && (input.Format == nil || *p.Format != *input.Format) {
		return false
	}

	if p.Formula != nil && (input.Formula == nil || *p.Formula != *input.Formula) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OutputType != nil && (input.OutputType == nil || *p.OutputType != *input.OutputType) {
		return false
	}

	return true
}
