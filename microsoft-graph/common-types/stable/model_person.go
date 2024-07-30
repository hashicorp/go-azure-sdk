package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Person struct {
	Birthday             *string               `json:"birthday,omitempty"`
	CompanyName          *string               `json:"companyName,omitempty"`
	Department           *string               `json:"department,omitempty"`
	DisplayName          *string               `json:"displayName,omitempty"`
	GivenName            *string               `json:"givenName,omitempty"`
	Id                   *string               `json:"id,omitempty"`
	ImAddress            *string               `json:"imAddress,omitempty"`
	IsFavorite           *bool                 `json:"isFavorite,omitempty"`
	JobTitle             *string               `json:"jobTitle,omitempty"`
	ODataType            *string               `json:"@odata.type,omitempty"`
	OfficeLocation       *string               `json:"officeLocation,omitempty"`
	PersonNotes          *string               `json:"personNotes,omitempty"`
	PersonType           *PersonType           `json:"personType,omitempty"`
	Phones               *[]Phone              `json:"phones,omitempty"`
	PostalAddresses      *[]Location           `json:"postalAddresses,omitempty"`
	Profession           *string               `json:"profession,omitempty"`
	ScoredEmailAddresses *[]ScoredEmailAddress `json:"scoredEmailAddresses,omitempty"`
	Surname              *string               `json:"surname,omitempty"`
	UserPrincipalName    *string               `json:"userPrincipalName,omitempty"`
	Websites             *[]Website            `json:"websites,omitempty"`
	YomiCompany          *string               `json:"yomiCompany,omitempty"`
}
