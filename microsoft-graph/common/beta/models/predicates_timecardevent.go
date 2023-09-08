package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardEventOperationPredicate struct {
	AtApprovedLocation *bool
	DateTime           *string
	ODataType          *string
}

func (p TimeCardEventOperationPredicate) Matches(input TimeCardEvent) bool {

	if p.AtApprovedLocation != nil && (input.AtApprovedLocation == nil || *p.AtApprovedLocation != *input.AtApprovedLocation) {
		return false
	}

	if p.DateTime != nil && (input.DateTime == nil || *p.DateTime != *input.DateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
