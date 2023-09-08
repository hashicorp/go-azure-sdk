package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSchool struct {
	Address              *PhysicalAddress               `json:"address,omitempty"`
	AdministrativeUnit   *AdministrativeUnit            `json:"administrativeUnit,omitempty"`
	Classes              *[]EducationClass              `json:"classes,omitempty"`
	CreatedBy            *IdentitySet                   `json:"createdBy,omitempty"`
	Description          *string                        `json:"description,omitempty"`
	DisplayName          *string                        `json:"displayName,omitempty"`
	ExternalId           *string                        `json:"externalId,omitempty"`
	ExternalPrincipalId  *string                        `json:"externalPrincipalId,omitempty"`
	ExternalSource       *EducationSchoolExternalSource `json:"externalSource,omitempty"`
	ExternalSourceDetail *string                        `json:"externalSourceDetail,omitempty"`
	Fax                  *string                        `json:"fax,omitempty"`
	HighestGrade         *string                        `json:"highestGrade,omitempty"`
	Id                   *string                        `json:"id,omitempty"`
	LowestGrade          *string                        `json:"lowestGrade,omitempty"`
	ODataType            *string                        `json:"@odata.type,omitempty"`
	Phone                *string                        `json:"phone,omitempty"`
	PrincipalEmail       *string                        `json:"principalEmail,omitempty"`
	PrincipalName        *string                        `json:"principalName,omitempty"`
	SchoolNumber         *string                        `json:"schoolNumber,omitempty"`
	Users                *[]EducationUser               `json:"users,omitempty"`
}
