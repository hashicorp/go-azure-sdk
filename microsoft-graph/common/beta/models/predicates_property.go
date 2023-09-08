package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyOperationPredicate struct {
	IsQueryable   *bool
	IsRefinable   *bool
	IsRetrievable *bool
	IsSearchable  *bool
	Name          *string
	ODataType     *string
}

func (p PropertyOperationPredicate) Matches(input Property) bool {

	if p.IsQueryable != nil && (input.IsQueryable == nil || *p.IsQueryable != *input.IsQueryable) {
		return false
	}

	if p.IsRefinable != nil && (input.IsRefinable == nil || *p.IsRefinable != *input.IsRefinable) {
		return false
	}

	if p.IsRetrievable != nil && (input.IsRetrievable == nil || *p.IsRetrievable != *input.IsRetrievable) {
		return false
	}

	if p.IsSearchable != nil && (input.IsSearchable == nil || *p.IsSearchable != *input.IsSearchable) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
