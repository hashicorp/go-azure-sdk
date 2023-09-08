package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountAliasOperationPredicate struct {
	Id        *string
	IdType    *string
	ODataType *string
}

func (p AccountAliasOperationPredicate) Matches(input AccountAlias) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IdType != nil && (input.IdType == nil || *p.IdType != *input.IdType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
