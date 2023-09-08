package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQueryStringOperationPredicate struct {
	ODataType *string
	Query     *string
}

func (p SearchQueryStringOperationPredicate) Matches(input SearchQueryString) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Query != nil && (input.Query == nil || *p.Query != *input.Query) {
		return false
	}

	return true
}
