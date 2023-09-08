package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummary struct {
	AssignmentFilterDisplayName              *string                                                    `json:"assignmentFilterDisplayName,omitempty"`
	AssignmentFilterId                       *string                                                    `json:"assignmentFilterId,omitempty"`
	AssignmentFilterLastModifiedDateTime     *string                                                    `json:"assignmentFilterLastModifiedDateTime,omitempty"`
	AssignmentFilterPlatform                 *AssignmentFilterEvaluationSummaryAssignmentFilterPlatform `json:"assignmentFilterPlatform,omitempty"`
	AssignmentFilterType                     *AssignmentFilterEvaluationSummaryAssignmentFilterType     `json:"assignmentFilterType,omitempty"`
	AssignmentFilterTypeAndEvaluationResults *[]AssignmentFilterTypeAndEvaluationResult                 `json:"assignmentFilterTypeAndEvaluationResults,omitempty"`
	EvaluationDateTime                       *string                                                    `json:"evaluationDateTime,omitempty"`
	EvaluationResult                         *AssignmentFilterEvaluationSummaryEvaluationResult         `json:"evaluationResult,omitempty"`
	ODataType                                *string                                                    `json:"@odata.type,omitempty"`
}
