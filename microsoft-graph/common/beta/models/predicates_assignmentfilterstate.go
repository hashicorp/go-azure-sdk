package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterStateOperationPredicate struct {
	Enabled   *bool
	ODataType *string
}

func (p AssignmentFilterStateOperationPredicate) Matches(input AssignmentFilterState) bool {

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
