package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResultInfoOperationPredicate struct {
	Code      *int64
	Message   *string
	ODataType *string
	Subcode   *int64
}

func (p ResultInfoOperationPredicate) Matches(input ResultInfo) bool {

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Subcode != nil && (input.Subcode == nil || *p.Subcode != *input.Subcode) {
		return false
	}

	return true
}
