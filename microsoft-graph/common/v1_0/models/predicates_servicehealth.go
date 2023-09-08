package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthOperationPredicate struct {
	Id        *string
	ODataType *string
	Service   *string
}

func (p ServiceHealthOperationPredicate) Matches(input ServiceHealth) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Service != nil && (input.Service == nil || *p.Service != *input.Service) {
		return false
	}

	return true
}
