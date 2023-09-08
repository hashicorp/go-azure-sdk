package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationClass struct {
	AssignmentCategories *[]EducationCategory          `json:"assignmentCategories,omitempty"`
	AssignmentDefaults   *EducationAssignmentDefaults  `json:"assignmentDefaults,omitempty"`
	AssignmentSettings   *EducationAssignmentSettings  `json:"assignmentSettings,omitempty"`
	Assignments          *[]EducationAssignment        `json:"assignments,omitempty"`
	ClassCode            *string                       `json:"classCode,omitempty"`
	Course               *EducationCourse              `json:"course,omitempty"`
	CreatedBy            *IdentitySet                  `json:"createdBy,omitempty"`
	Description          *string                       `json:"description,omitempty"`
	DisplayName          *string                       `json:"displayName,omitempty"`
	ExternalId           *string                       `json:"externalId,omitempty"`
	ExternalName         *string                       `json:"externalName,omitempty"`
	ExternalSource       *EducationClassExternalSource `json:"externalSource,omitempty"`
	ExternalSourceDetail *string                       `json:"externalSourceDetail,omitempty"`
	Grade                *string                       `json:"grade,omitempty"`
	Group                *Group                        `json:"group,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	MailNickname         *string                       `json:"mailNickname,omitempty"`
	Members              *[]EducationUser              `json:"members,omitempty"`
	Modules              *[]EducationModule            `json:"modules,omitempty"`
	ODataType            *string                       `json:"@odata.type,omitempty"`
	Schools              *[]EducationSchool            `json:"schools,omitempty"`
	Teachers             *[]EducationUser              `json:"teachers,omitempty"`
	Term                 *EducationTerm                `json:"term,omitempty"`
}
