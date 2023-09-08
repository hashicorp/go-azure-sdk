package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkHostedContentOperationPredicate struct {
	ContentBytes *string
	ContentType  *string
	Id           *string
	ODataType    *string
}

func (p TeamworkHostedContentOperationPredicate) Matches(input TeamworkHostedContent) bool {

	if p.ContentBytes != nil && (input.ContentBytes == nil || *p.ContentBytes != *input.ContentBytes) {
		return false
	}

	if p.ContentType != nil && (input.ContentType == nil || *p.ContentType != *input.ContentType) {
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
