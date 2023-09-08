package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingStatusInfoOperationPredicate struct {
	AssignedDateTime   *string
	CompletionDateTime *string
	DisplayName        *string
	ODataType          *string
}

func (p UserTrainingStatusInfoOperationPredicate) Matches(input UserTrainingStatusInfo) bool {

	if p.AssignedDateTime != nil && (input.AssignedDateTime == nil || *p.AssignedDateTime != *input.AssignedDateTime) {
		return false
	}

	if p.CompletionDateTime != nil && (input.CompletionDateTime == nil || *p.CompletionDateTime != *input.CompletionDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
