package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskOperationPredicate struct {
	Id        *string
	ODataType *string
	ParentUrl *string
}

func (p PrintTaskOperationPredicate) Matches(input PrintTask) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentUrl != nil && (input.ParentUrl == nil || *p.ParentUrl != *input.ParentUrl) {
		return false
	}

	return true
}
