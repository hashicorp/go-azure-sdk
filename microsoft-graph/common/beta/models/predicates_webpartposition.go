package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebPartPositionOperationPredicate struct {
	IsInVerticalSection *bool
	ODataType           *string
}

func (p WebPartPositionOperationPredicate) Matches(input WebPartPosition) bool {

	if p.IsInVerticalSection != nil && (input.IsInVerticalSection == nil || *p.IsInVerticalSection != *input.IsInVerticalSection) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
