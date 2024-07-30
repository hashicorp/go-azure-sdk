package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantDetailedInformation struct {
	City              *string `json:"city,omitempty"`
	CountryCode       *string `json:"countryCode,omitempty"`
	CountryName       *string `json:"countryName,omitempty"`
	DefaultDomainName *string `json:"defaultDomainName,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	Id                *string `json:"id,omitempty"`
	IndustryName      *string `json:"industryName,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	Region            *string `json:"region,omitempty"`
	SegmentName       *string `json:"segmentName,omitempty"`
	TenantId          *string `json:"tenantId,omitempty"`
	VerticalName      *string `json:"verticalName,omitempty"`
}
