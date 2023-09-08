package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccountInformationOperationPredicate struct {
	AgeGroup             *string
	CountryCode          *string
	CreatedDateTime      *string
	Id                   *string
	IsSearchable         *bool
	LastModifiedDateTime *string
	ODataType            *string
	UserPrincipalName    *string
}

func (p UserAccountInformationOperationPredicate) Matches(input UserAccountInformation) bool {

	if p.AgeGroup != nil && (input.AgeGroup == nil || *p.AgeGroup != *input.AgeGroup) {
		return false
	}

	if p.CountryCode != nil && (input.CountryCode == nil || *p.CountryCode != *input.CountryCode) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSearchable != nil && (input.IsSearchable == nil || *p.IsSearchable != *input.IsSearchable) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
