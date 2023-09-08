package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrgContactOperationPredicate struct {
	CompanyName                *string
	DeletedDateTime            *string
	Department                 *string
	DisplayName                *string
	GivenName                  *string
	Id                         *string
	JobTitle                   *string
	Mail                       *string
	MailNickname               *string
	ODataType                  *string
	OnPremisesLastSyncDateTime *string
	OnPremisesSyncEnabled      *bool
	Surname                    *string
}

func (p OrgContactOperationPredicate) Matches(input OrgContact) bool {

	if p.CompanyName != nil && (input.CompanyName == nil || *p.CompanyName != *input.CompanyName) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
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

	if p.JobTitle != nil && (input.JobTitle == nil || *p.JobTitle != *input.JobTitle) {
		return false
	}

	if p.Mail != nil && (input.Mail == nil || *p.Mail != *input.Mail) {
		return false
	}

	if p.MailNickname != nil && (input.MailNickname == nil || *p.MailNickname != *input.MailNickname) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnPremisesLastSyncDateTime != nil && (input.OnPremisesLastSyncDateTime == nil || *p.OnPremisesLastSyncDateTime != *input.OnPremisesLastSyncDateTime) {
		return false
	}

	if p.OnPremisesSyncEnabled != nil && (input.OnPremisesSyncEnabled == nil || *p.OnPremisesSyncEnabled != *input.OnPremisesSyncEnabled) {
		return false
	}

	if p.Surname != nil && (input.Surname == nil || *p.Surname != *input.Surname) {
		return false
	}

	return true
}
