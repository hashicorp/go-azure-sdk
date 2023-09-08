package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Employee struct {
	Address              *PostalAddressType `json:"address,omitempty"`
	BirthDate            *string            `json:"birthDate,omitempty"`
	DisplayName          *string            `json:"displayName,omitempty"`
	Email                *string            `json:"email,omitempty"`
	EmploymentDate       *string            `json:"employmentDate,omitempty"`
	GivenName            *string            `json:"givenName,omitempty"`
	Id                   *string            `json:"id,omitempty"`
	JobTitle             *string            `json:"jobTitle,omitempty"`
	LastModifiedDateTime *string            `json:"lastModifiedDateTime,omitempty"`
	MiddleName           *string            `json:"middleName,omitempty"`
	MobilePhone          *string            `json:"mobilePhone,omitempty"`
	Number               *string            `json:"number,omitempty"`
	ODataType            *string            `json:"@odata.type,omitempty"`
	PersonalEmail        *string            `json:"personalEmail,omitempty"`
	PhoneNumber          *string            `json:"phoneNumber,omitempty"`
	Picture              *[]Picture         `json:"picture,omitempty"`
	StatisticsGroupCode  *string            `json:"statisticsGroupCode,omitempty"`
	Status               *string            `json:"status,omitempty"`
	Surname              *string            `json:"surname,omitempty"`
	TerminationDate      *string            `json:"terminationDate,omitempty"`
}
