package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DateTimeColumnOperationPredicate struct {
	DisplayAs *string
	Format    *string
	ODataType *string
}

func (p DateTimeColumnOperationPredicate) Matches(input DateTimeColumn) bool {

	if p.DisplayAs != nil && (input.DisplayAs == nil || *p.DisplayAs != *input.DisplayAs) {
		return false
	}

	if p.Format != nil && (input.Format == nil || *p.Format != *input.Format) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
