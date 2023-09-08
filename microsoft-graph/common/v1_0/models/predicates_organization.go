package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationOperationPredicate struct {
	City                       *string
	Country                    *string
	CountryLetterCode          *string
	CreatedDateTime            *string
	DefaultUsageLocation       *string
	DeletedDateTime            *string
	DisplayName                *string
	Id                         *string
	ODataType                  *string
	OnPremisesLastSyncDateTime *string
	OnPremisesSyncEnabled      *bool
	PostalCode                 *string
	PreferredLanguage          *string
	State                      *string
	Street                     *string
	TenantType                 *string
}

func (p OrganizationOperationPredicate) Matches(input Organization) bool {

	if p.City != nil && (input.City == nil || *p.City != *input.City) {
		return false
	}

	if p.Country != nil && (input.Country == nil || *p.Country != *input.Country) {
		return false
	}

	if p.CountryLetterCode != nil && (input.CountryLetterCode == nil || *p.CountryLetterCode != *input.CountryLetterCode) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DefaultUsageLocation != nil && (input.DefaultUsageLocation == nil || *p.DefaultUsageLocation != *input.DefaultUsageLocation) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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

	if p.PostalCode != nil && (input.PostalCode == nil || *p.PostalCode != *input.PostalCode) {
		return false
	}

	if p.PreferredLanguage != nil && (input.PreferredLanguage == nil || *p.PreferredLanguage != *input.PreferredLanguage) {
		return false
	}

	if p.State != nil && (input.State == nil || *p.State != *input.State) {
		return false
	}

	if p.Street != nil && (input.Street == nil || *p.Street != *input.Street) {
		return false
	}

	if p.TenantType != nil && (input.TenantType == nil || *p.TenantType != *input.TenantType) {
		return false
	}

	return true
}
