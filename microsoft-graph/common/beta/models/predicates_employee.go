package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeOperationPredicate struct {
	BirthDate            *string
	DisplayName          *string
	Email                *string
	EmploymentDate       *string
	GivenName            *string
	Id                   *string
	JobTitle             *string
	LastModifiedDateTime *string
	MiddleName           *string
	MobilePhone          *string
	Number               *string
	ODataType            *string
	PersonalEmail        *string
	PhoneNumber          *string
	StatisticsGroupCode  *string
	Status               *string
	Surname              *string
	TerminationDate      *string
}

func (p EmployeeOperationPredicate) Matches(input Employee) bool {

	if p.BirthDate != nil && (input.BirthDate == nil || *p.BirthDate != *input.BirthDate) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.EmploymentDate != nil && (input.EmploymentDate == nil || *p.EmploymentDate != *input.EmploymentDate) {
		return false
	}

	if p.GivenName != nil && (input.GivenName == nil || *p.GivenName != *input.GivenName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.JobTitle != nil && (input.JobTitle == nil || *p.JobTitle != *input.JobTitle) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MiddleName != nil && (input.MiddleName == nil || *p.MiddleName != *input.MiddleName) {
		return false
	}

	if p.MobilePhone != nil && (input.MobilePhone == nil || *p.MobilePhone != *input.MobilePhone) {
		return false
	}

	if p.Number != nil && (input.Number == nil || *p.Number != *input.Number) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PersonalEmail != nil && (input.PersonalEmail == nil || *p.PersonalEmail != *input.PersonalEmail) {
		return false
	}

	if p.PhoneNumber != nil && (input.PhoneNumber == nil || *p.PhoneNumber != *input.PhoneNumber) {
		return false
	}

	if p.StatisticsGroupCode != nil && (input.StatisticsGroupCode == nil || *p.StatisticsGroupCode != *input.StatisticsGroupCode) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	if p.Surname != nil && (input.Surname == nil || *p.Surname != *input.Surname) {
		return false
	}

	if p.TerminationDate != nil && (input.TerminationDate == nil || *p.TerminationDate != *input.TerminationDate) {
		return false
	}

	return true
}
