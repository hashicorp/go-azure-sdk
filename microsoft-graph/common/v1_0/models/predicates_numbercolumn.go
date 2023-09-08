package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumberColumnOperationPredicate struct {
	DecimalPlaces *string
	DisplayAs     *string
	ODataType     *string
}

func (p NumberColumnOperationPredicate) Matches(input NumberColumn) bool {

	if p.DecimalPlaces != nil && (input.DecimalPlaces == nil || *p.DecimalPlaces != *input.DecimalPlaces) {
		return false
	}

	if p.DisplayAs != nil && (input.DisplayAs == nil || *p.DisplayAs != *input.DisplayAs) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
