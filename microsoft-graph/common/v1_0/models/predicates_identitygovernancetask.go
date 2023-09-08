package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskOperationPredicate struct {
	ContinueOnError   *bool
	Description       *string
	DisplayName       *string
	ExecutionSequence *int64
	Id                *string
	IsEnabled         *bool
	ODataType         *string
	TaskDefinitionId  *string
}

func (p IdentityGovernanceTaskOperationPredicate) Matches(input IdentityGovernanceTask) bool {

	if p.ContinueOnError != nil && (input.ContinueOnError == nil || *p.ContinueOnError != *input.ContinueOnError) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExecutionSequence != nil && (input.ExecutionSequence == nil || *p.ExecutionSequence != *input.ExecutionSequence) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TaskDefinitionId != nil && (input.TaskDefinitionId == nil || *p.TaskDefinitionId != *input.TaskDefinitionId) {
		return false
	}

	return true
}
