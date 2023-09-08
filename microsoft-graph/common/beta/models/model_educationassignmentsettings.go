package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentSettings struct {
	GradingCategories           *[]EducationGradingCategory `json:"gradingCategories,omitempty"`
	Id                          *string                     `json:"id,omitempty"`
	ODataType                   *string                     `json:"@odata.type,omitempty"`
	SubmissionAnimationDisabled *bool                       `json:"submissionAnimationDisabled,omitempty"`
}
