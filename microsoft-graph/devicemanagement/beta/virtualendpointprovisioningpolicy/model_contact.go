package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Contact struct {
	AssistantName                 *string                              `json:"assistantName,omitempty"`
	Birthday                      *string                              `json:"birthday,omitempty"`
	Categories                    *[]string                            `json:"categories,omitempty"`
	ChangeKey                     *string                              `json:"changeKey,omitempty"`
	Children                      *[]string                            `json:"children,omitempty"`
	CompanyName                   *string                              `json:"companyName,omitempty"`
	CreatedDateTime               *string                              `json:"createdDateTime,omitempty"`
	Department                    *string                              `json:"department,omitempty"`
	DisplayName                   *string                              `json:"displayName,omitempty"`
	EmailAddresses                *[]TypedEmailAddress                 `json:"emailAddresses,omitempty"`
	Extensions                    *[]Extension                         `json:"extensions,omitempty"`
	FileAs                        *string                              `json:"fileAs,omitempty"`
	Flag                          *FollowupFlag                        `json:"flag,omitempty"`
	Gender                        *string                              `json:"gender,omitempty"`
	Generation                    *string                              `json:"generation,omitempty"`
	GivenName                     *string                              `json:"givenName,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	ImAddresses                   *[]string                            `json:"imAddresses,omitempty"`
	Initials                      *string                              `json:"initials,omitempty"`
	IsFavorite                    *bool                                `json:"isFavorite,omitempty"`
	JobTitle                      *string                              `json:"jobTitle,omitempty"`
	LastModifiedDateTime          *string                              `json:"lastModifiedDateTime,omitempty"`
	Manager                       *string                              `json:"manager,omitempty"`
	MiddleName                    *string                              `json:"middleName,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	NickName                      *string                              `json:"nickName,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	OfficeLocation                *string                              `json:"officeLocation,omitempty"`
	ParentFolderId                *string                              `json:"parentFolderId,omitempty"`
	PersonalNotes                 *string                              `json:"personalNotes,omitempty"`
	Phones                        *[]Phone                             `json:"phones,omitempty"`
	Photo                         *ProfilePhoto                        `json:"photo,omitempty"`
	PostalAddresses               *[]PhysicalAddress                   `json:"postalAddresses,omitempty"`
	Profession                    *string                              `json:"profession,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	SpouseName                    *string                              `json:"spouseName,omitempty"`
	Surname                       *string                              `json:"surname,omitempty"`
	Title                         *string                              `json:"title,omitempty"`
	Websites                      *[]Website                           `json:"websites,omitempty"`
	WeddingAnniversary            *string                              `json:"weddingAnniversary,omitempty"`
	YomiCompanyName               *string                              `json:"yomiCompanyName,omitempty"`
	YomiGivenName                 *string                              `json:"yomiGivenName,omitempty"`
	YomiSurname                   *string                              `json:"yomiSurname,omitempty"`
}
