package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationStudent struct {
	BirthDate      *string                 `json:"birthDate,omitempty"`
	ExternalId     *string                 `json:"externalId,omitempty"`
	Gender         *EducationStudentGender `json:"gender,omitempty"`
	Grade          *string                 `json:"grade,omitempty"`
	GraduationYear *string                 `json:"graduationYear,omitempty"`
	ODataType      *string                 `json:"@odata.type,omitempty"`
	StudentNumber  *string                 `json:"studentNumber,omitempty"`
}
