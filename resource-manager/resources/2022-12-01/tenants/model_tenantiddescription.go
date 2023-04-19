package tenants

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantIdDescription struct {
	Country               *string         `json:"country,omitempty"`
	CountryCode           *string         `json:"countryCode,omitempty"`
	DefaultDomain         *string         `json:"defaultDomain,omitempty"`
	DisplayName           *string         `json:"displayName,omitempty"`
	Domains               *[]string       `json:"domains,omitempty"`
	Id                    *string         `json:"id,omitempty"`
	TenantBrandingLogoUrl *string         `json:"tenantBrandingLogoUrl,omitempty"`
	TenantCategory        *TenantCategory `json:"tenantCategory,omitempty"`
	TenantId              *string         `json:"tenantId,omitempty"`
	TenantType            *string         `json:"tenantType,omitempty"`
}
