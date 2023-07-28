package tenants

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantIdDescriptionOperationPredicate struct {
	Country               *string
	CountryCode           *string
	DefaultDomain         *string
	DisplayName           *string
	Id                    *string
	TenantBrandingLogoUrl *string
	TenantId              *string
	TenantType            *string
}

func (p TenantIdDescriptionOperationPredicate) Matches(input TenantIdDescription) bool {

	if p.Country != nil && (input.Country == nil || *p.Country != *input.Country) {
		return false
	}

	if p.CountryCode != nil && (input.CountryCode == nil || *p.CountryCode != *input.CountryCode) {
		return false
	}

	if p.DefaultDomain != nil && (input.DefaultDomain == nil || *p.DefaultDomain != *input.DefaultDomain) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.TenantBrandingLogoUrl != nil && (input.TenantBrandingLogoUrl == nil || *p.TenantBrandingLogoUrl != *input.TenantBrandingLogoUrl) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.TenantType != nil && (input.TenantType == nil || *p.TenantType != *input.TenantType) {
		return false
	}

	return true
}
