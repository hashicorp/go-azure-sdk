package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesFilterOperationPredicate struct {
	ODataType *string
	Rule      *string
}

func (p DevicesFilterOperationPredicate) Matches(input DevicesFilter) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Rule != nil && (input.Rule == nil || *p.Rule != *input.Rule) {
		return false
	}

	return true
}
