package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonOperationPredicate struct {
	Birthday          *string
	CompanyName       *string
	Department        *string
	DisplayName       *string
	GivenName         *string
	Id                *string
	ImAddress         *string
	IsFavorite        *bool
	JobTitle          *string
	ODataType         *string
	OfficeLocation    *string
	PersonNotes       *string
	Profession        *string
	Surname           *string
	UserPrincipalName *string
	YomiCompany       *string
}

func (p PersonOperationPredicate) Matches(input Person) bool {

	if p.Birthday != nil && (input.Birthday == nil || *p.Birthday != *input.Birthday) {
		return false
	}

	if p.CompanyName != nil && (input.CompanyName == nil || *p.CompanyName != *input.CompanyName) {
		return false
	}

	if p.Department != nil && (input.Department == nil || *p.Department != *input.Department) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.GivenName != nil && (input.GivenName == nil || *p.GivenName != *input.GivenName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ImAddress != nil && (input.ImAddress == nil || *p.ImAddress != *input.ImAddress) {
		return false
	}

	if p.IsFavorite != nil && (input.IsFavorite == nil || *p.IsFavorite != *input.IsFavorite) {
		return false
	}

	if p.JobTitle != nil && (input.JobTitle == nil || *p.JobTitle != *input.JobTitle) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OfficeLocation != nil && (input.OfficeLocation == nil || *p.OfficeLocation != *input.OfficeLocation) {
		return false
	}

	if p.PersonNotes != nil && (input.PersonNotes == nil || *p.PersonNotes != *input.PersonNotes) {
		return false
	}

	if p.Profession != nil && (input.Profession == nil || *p.Profession != *input.Profession) {
		return false
	}

	if p.Surname != nil && (input.Surname == nil || *p.Surname != *input.Surname) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	if p.YomiCompany != nil && (input.YomiCompany == nil || *p.YomiCompany != *input.YomiCompany) {
		return false
	}

	return true
}
