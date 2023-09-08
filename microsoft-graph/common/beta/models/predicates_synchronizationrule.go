package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationRuleOperationPredicate struct {
	Editable            *bool
	Id                  *string
	Name                *string
	ODataType           *string
	Priority            *int64
	SourceDirectoryName *string
	TargetDirectoryName *string
}

func (p SynchronizationRuleOperationPredicate) Matches(input SynchronizationRule) bool {

	if p.Editable != nil && (input.Editable == nil || *p.Editable != *input.Editable) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	if p.SourceDirectoryName != nil && (input.SourceDirectoryName == nil || *p.SourceDirectoryName != *input.SourceDirectoryName) {
		return false
	}

	if p.TargetDirectoryName != nil && (input.TargetDirectoryName == nil || *p.TargetDirectoryName != *input.TargetDirectoryName) {
		return false
	}

	return true
}
