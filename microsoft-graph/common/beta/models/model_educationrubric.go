package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationRubric struct {
	CreatedBy            *IdentitySet                  `json:"createdBy,omitempty"`
	CreatedDateTime      *string                       `json:"createdDateTime,omitempty"`
	Description          *EducationItemBody            `json:"description,omitempty"`
	DisplayName          *string                       `json:"displayName,omitempty"`
	Grading              *EducationAssignmentGradeType `json:"grading,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet                  `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                       `json:"lastModifiedDateTime,omitempty"`
	Levels               *[]RubricLevel                `json:"levels,omitempty"`
	ODataType            *string                       `json:"@odata.type,omitempty"`
	Qualities            *[]RubricQuality              `json:"qualities,omitempty"`
}
