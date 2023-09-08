package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionRequestOperationPredicate struct {
	CompletedDateTime *string
	ErrorMessage      *string
	Id                *string
	ODataType         *string
}

func (p AppLogCollectionRequestOperationPredicate) Matches(input AppLogCollectionRequest) bool {

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.ErrorMessage != nil && (input.ErrorMessage == nil || *p.ErrorMessage != *input.ErrorMessage) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
