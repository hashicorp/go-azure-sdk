package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionOperationPredicate struct {
	ActionWebUrl *string
	ODataType    *string
	Title        *string
}

func (p RecommendedActionOperationPredicate) Matches(input RecommendedAction) bool {

	if p.ActionWebUrl != nil && (input.ActionWebUrl == nil || *p.ActionWebUrl != *input.ActionWebUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
