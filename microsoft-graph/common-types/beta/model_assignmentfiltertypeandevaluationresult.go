package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterTypeAndEvaluationResult struct {
	AssignmentFilterType *AssignmentFilterTypeAndEvaluationResultAssignmentFilterType `json:"assignmentFilterType,omitempty"`
	EvaluationResult     *AssignmentFilterTypeAndEvaluationResultEvaluationResult     `json:"evaluationResult,omitempty"`
	ODataType            *string                                                      `json:"@odata.type,omitempty"`
}
