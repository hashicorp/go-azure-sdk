package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricLevel struct {
	Description *EducationItemBody            `json:"description,omitempty"`
	DisplayName *string                       `json:"displayName,omitempty"`
	Grading     *EducationAssignmentGradeType `json:"grading,omitempty"`
	LevelId     *string                       `json:"levelId,omitempty"`
	ODataType   *string                       `json:"@odata.type,omitempty"`
}
