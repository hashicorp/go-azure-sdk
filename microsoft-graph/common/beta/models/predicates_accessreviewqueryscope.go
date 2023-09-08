package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewQueryScopeOperationPredicate struct {
	ODataType *string
	Query     *string
	QueryRoot *string
	QueryType *string
}

func (p AccessReviewQueryScopeOperationPredicate) Matches(input AccessReviewQueryScope) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Query != nil && (input.Query == nil || *p.Query != *input.Query) {
		return false
	}

	if p.QueryRoot != nil && (input.QueryRoot == nil || *p.QueryRoot != *input.QueryRoot) {
		return false
	}

	if p.QueryType != nil && (input.QueryType == nil || *p.QueryType != *input.QueryType) {
		return false
	}

	return true
}
