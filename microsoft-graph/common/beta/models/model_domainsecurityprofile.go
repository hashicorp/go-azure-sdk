package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainSecurityProfile struct {
	ActivityGroupNames       *[]string                  `json:"activityGroupNames,omitempty"`
	AzureSubscriptionId      *string                    `json:"azureSubscriptionId,omitempty"`
	AzureTenantId            *string                    `json:"azureTenantId,omitempty"`
	CountHits                *int64                     `json:"countHits,omitempty"`
	CountInOrg               *int64                     `json:"countInOrg,omitempty"`
	DomainCategories         *[]ReputationCategory      `json:"domainCategories,omitempty"`
	DomainRegisteredDateTime *string                    `json:"domainRegisteredDateTime,omitempty"`
	FirstSeenDateTime        *string                    `json:"firstSeenDateTime,omitempty"`
	Id                       *string                    `json:"id,omitempty"`
	LastSeenDateTime         *string                    `json:"lastSeenDateTime,omitempty"`
	Name                     *string                    `json:"name,omitempty"`
	ODataType                *string                    `json:"@odata.type,omitempty"`
	Registrant               *DomainRegistrant          `json:"registrant,omitempty"`
	RiskScore                *string                    `json:"riskScore,omitempty"`
	Tags                     *[]string                  `json:"tags,omitempty"`
	VendorInformation        *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}
