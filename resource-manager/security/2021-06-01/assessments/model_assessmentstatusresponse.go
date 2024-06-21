package assessments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentStatusResponse struct {
	Cause               *string              `json:"cause,omitempty"`
	Code                AssessmentStatusCode `json:"code"`
	Description         *string              `json:"description,omitempty"`
	FirstEvaluationDate *string              `json:"firstEvaluationDate,omitempty"`
	StatusChangeDate    *string              `json:"statusChangeDate,omitempty"`
}
