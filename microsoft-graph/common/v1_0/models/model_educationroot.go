package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationRoot struct {
	Classes   *[]EducationClass  `json:"classes,omitempty"`
	Me        *EducationUser     `json:"me,omitempty"`
	ODataType *string            `json:"@odata.type,omitempty"`
	Schools   *[]EducationSchool `json:"schools,omitempty"`
	Users     *[]EducationUser   `json:"users,omitempty"`
}
