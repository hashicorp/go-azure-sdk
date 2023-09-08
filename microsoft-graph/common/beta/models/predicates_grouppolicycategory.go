package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryOperationPredicate struct {
	DisplayName          *string
	Id                   *string
	IsRoot               *bool
	LastModifiedDateTime *string
	ODataType            *string
}

func (p GroupPolicyCategoryOperationPredicate) Matches(input GroupPolicyCategory) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsRoot != nil && (input.IsRoot == nil || *p.IsRoot != *input.IsRoot) {
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
