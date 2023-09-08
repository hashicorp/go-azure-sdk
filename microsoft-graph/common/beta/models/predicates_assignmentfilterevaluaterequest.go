package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluateRequestOperationPredicate struct {
	ODataType *string
	Rule      *string
	Search    *string
	Skip      *int64
	Top       *int64
}

func (p AssignmentFilterEvaluateRequestOperationPredicate) Matches(input AssignmentFilterEvaluateRequest) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Rule != nil && (input.Rule == nil || *p.Rule != *input.Rule) {
		return false
	}

	if p.Search != nil && (input.Search == nil || *p.Search != *input.Search) {
		return false
	}

	if p.Skip != nil && (input.Skip == nil || *p.Skip != *input.Skip) {
		return false
	}

	if p.Top != nil && (input.Top == nil || *p.Top != *input.Top) {
		return false
	}

	return true
}
