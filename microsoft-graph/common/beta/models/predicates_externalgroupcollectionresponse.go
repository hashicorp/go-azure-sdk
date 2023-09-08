package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalGroupCollectionResponseOperationPredicate struct {
	ODataCount    *int64
	ODataNextLink *string
}

func (p ExternalGroupCollectionResponseOperationPredicate) Matches(input ExternalGroupCollectionResponse) bool {

	if p.ODataCount != nil && (input.ODataCount == nil || *p.ODataCount != *input.ODataCount) {
		return false
	}

	if p.ODataNextLink != nil && (input.ODataNextLink == nil || *p.ODataNextLink != *input.ODataNextLink) {
		return false
	}

	return true
}
