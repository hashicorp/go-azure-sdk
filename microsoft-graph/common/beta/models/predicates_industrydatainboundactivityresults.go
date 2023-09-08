package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundActivityResultsOperationPredicate struct {
	ActivityId  *string
	DisplayName *string
	Errors      *int64
	ODataType   *string
	Warnings    *int64
}

func (p IndustryDataInboundActivityResultsOperationPredicate) Matches(input IndustryDataInboundActivityResults) bool {

	if p.ActivityId != nil && (input.ActivityId == nil || *p.ActivityId != *input.ActivityId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Errors != nil && (input.Errors == nil || *p.Errors != *input.Errors) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Warnings != nil && (input.Warnings == nil || *p.Warnings != *input.Warnings) {
		return false
	}

	return true
}
