package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertTemplateOperationPredicate struct {
	Category           *string
	Description        *string
	ODataType          *string
	RecommendedActions *string
	Title              *string
}

func (p AlertTemplateOperationPredicate) Matches(input AlertTemplate) bool {

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecommendedActions != nil && (input.RecommendedActions == nil || *p.RecommendedActions != *input.RecommendedActions) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
