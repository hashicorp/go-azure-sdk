package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesResourceConnectionOperationPredicate struct {
	Id        *string
	ODataType *string
}

func (p WindowsUpdatesResourceConnectionOperationPredicate) Matches(input WindowsUpdatesResourceConnection) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
