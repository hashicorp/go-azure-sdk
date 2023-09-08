package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummaryOperationPredicate struct {
	AssignmentFilterDisplayName          *string
	AssignmentFilterId                   *string
	AssignmentFilterLastModifiedDateTime *string
	EvaluationDateTime                   *string
	ODataType                            *string
}

func (p AssignmentFilterEvaluationSummaryOperationPredicate) Matches(input AssignmentFilterEvaluationSummary) bool {

	if p.AssignmentFilterDisplayName != nil && (input.AssignmentFilterDisplayName == nil || *p.AssignmentFilterDisplayName != *input.AssignmentFilterDisplayName) {
		return false
	}

	if p.AssignmentFilterId != nil && (input.AssignmentFilterId == nil || *p.AssignmentFilterId != *input.AssignmentFilterId) {
		return false
	}

	if p.AssignmentFilterLastModifiedDateTime != nil && (input.AssignmentFilterLastModifiedDateTime == nil || *p.AssignmentFilterLastModifiedDateTime != *input.AssignmentFilterLastModifiedDateTime) {
		return false
	}

	if p.EvaluationDateTime != nil && (input.EvaluationDateTime == nil || *p.EvaluationDateTime != *input.EvaluationDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
