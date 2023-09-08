package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadOperationPredicate struct {
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	IsAutomated          *bool
	IsControversial      *bool
	IsCurrentEvent       *bool
	Language             *string
	LastModifiedDateTime *string
	ODataType            *string
}

func (p PayloadOperationPredicate) Matches(input Payload) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAutomated != nil && (input.IsAutomated == nil || *p.IsAutomated != *input.IsAutomated) {
		return false
	}

	if p.IsControversial != nil && (input.IsControversial == nil || *p.IsControversial != *input.IsControversial) {
		return false
	}

	if p.IsCurrentEvent != nil && (input.IsCurrentEvent == nil || *p.IsCurrentEvent != *input.IsCurrentEvent) {
		return false
	}

	if p.Language != nil && (input.Language == nil || *p.Language != *input.Language) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
