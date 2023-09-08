package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterValidationResultOperationPredicate struct {
	IsValidRule *bool
	ODataType   *string
}

func (p AssignmentFilterValidationResultOperationPredicate) Matches(input AssignmentFilterValidationResult) bool {

	if p.IsValidRule != nil && (input.IsValidRule == nil || *p.IsValidRule != *input.IsValidRule) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
