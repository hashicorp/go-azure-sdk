package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationClassOperationPredicate struct {
	ClassCode            *string
	Description          *string
	DisplayName          *string
	ExternalId           *string
	ExternalName         *string
	ExternalSourceDetail *string
	Grade                *string
	Id                   *string
	MailNickname         *string
	ODataType            *string
}

func (p EducationClassOperationPredicate) Matches(input EducationClass) bool {

	if p.ClassCode != nil && (input.ClassCode == nil || *p.ClassCode != *input.ClassCode) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.ExternalName != nil && (input.ExternalName == nil || *p.ExternalName != *input.ExternalName) {
		return false
	}

	if p.ExternalSourceDetail != nil && (input.ExternalSourceDetail == nil || *p.ExternalSourceDetail != *input.ExternalSourceDetail) {
		return false
	}

	if p.Grade != nil && (input.Grade == nil || *p.Grade != *input.Grade) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MailNickname != nil && (input.MailNickname == nil || *p.MailNickname != *input.MailNickname) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
