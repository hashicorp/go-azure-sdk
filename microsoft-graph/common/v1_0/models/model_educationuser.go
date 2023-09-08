package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationUser struct {
	AccountEnabled                 *bool                        `json:"accountEnabled,omitempty"`
	AssignedLicenses               *[]AssignedLicense           `json:"assignedLicenses,omitempty"`
	AssignedPlans                  *[]AssignedPlan              `json:"assignedPlans,omitempty"`
	Assignments                    *[]EducationAssignment       `json:"assignments,omitempty"`
	BusinessPhones                 *[]string                    `json:"businessPhones,omitempty"`
	Classes                        *[]EducationClass            `json:"classes,omitempty"`
	CreatedBy                      *IdentitySet                 `json:"createdBy,omitempty"`
	Department                     *string                      `json:"department,omitempty"`
	DisplayName                    *string                      `json:"displayName,omitempty"`
	ExternalSource                 *EducationUserExternalSource `json:"externalSource,omitempty"`
	ExternalSourceDetail           *string                      `json:"externalSourceDetail,omitempty"`
	GivenName                      *string                      `json:"givenName,omitempty"`
	Id                             *string                      `json:"id,omitempty"`
	Mail                           *string                      `json:"mail,omitempty"`
	MailNickname                   *string                      `json:"mailNickname,omitempty"`
	MailingAddress                 *PhysicalAddress             `json:"mailingAddress,omitempty"`
	MiddleName                     *string                      `json:"middleName,omitempty"`
	MobilePhone                    *string                      `json:"mobilePhone,omitempty"`
	ODataType                      *string                      `json:"@odata.type,omitempty"`
	OfficeLocation                 *string                      `json:"officeLocation,omitempty"`
	OnPremisesInfo                 *EducationOnPremisesInfo     `json:"onPremisesInfo,omitempty"`
	PasswordPolicies               *string                      `json:"passwordPolicies,omitempty"`
	PasswordProfile                *PasswordProfile             `json:"passwordProfile,omitempty"`
	PreferredLanguage              *string                      `json:"preferredLanguage,omitempty"`
	PrimaryRole                    *EducationUserPrimaryRole    `json:"primaryRole,omitempty"`
	ProvisionedPlans               *[]ProvisionedPlan           `json:"provisionedPlans,omitempty"`
	RefreshTokensValidFromDateTime *string                      `json:"refreshTokensValidFromDateTime,omitempty"`
	RelatedContacts                *[]RelatedContact            `json:"relatedContacts,omitempty"`
	ResidenceAddress               *PhysicalAddress             `json:"residenceAddress,omitempty"`
	Rubrics                        *[]EducationRubric           `json:"rubrics,omitempty"`
	Schools                        *[]EducationSchool           `json:"schools,omitempty"`
	ShowInAddressList              *bool                        `json:"showInAddressList,omitempty"`
	Student                        *EducationStudent            `json:"student,omitempty"`
	Surname                        *string                      `json:"surname,omitempty"`
	TaughtClasses                  *[]EducationClass            `json:"taughtClasses,omitempty"`
	Teacher                        *EducationTeacher            `json:"teacher,omitempty"`
	UsageLocation                  *string                      `json:"usageLocation,omitempty"`
	User                           *User                        `json:"user,omitempty"`
	UserPrincipalName              *string                      `json:"userPrincipalName,omitempty"`
	UserType                       *string                      `json:"userType,omitempty"`
}
