package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostSecurityProfile struct {
	AzureSubscriptionId       *string                    `json:"azureSubscriptionId,omitempty"`
	AzureTenantId             *string                    `json:"azureTenantId,omitempty"`
	FirstSeenDateTime         *string                    `json:"firstSeenDateTime,omitempty"`
	Fqdn                      *string                    `json:"fqdn,omitempty"`
	Id                        *string                    `json:"id,omitempty"`
	IsAzureAdJoined           *bool                      `json:"isAzureAdJoined,omitempty"`
	IsAzureAdRegistered       *bool                      `json:"isAzureAdRegistered,omitempty"`
	IsHybridAzureDomainJoined *bool                      `json:"isHybridAzureDomainJoined,omitempty"`
	LastSeenDateTime          *string                    `json:"lastSeenDateTime,omitempty"`
	LogonUsers                *[]LogonUser               `json:"logonUsers,omitempty"`
	NetBiosName               *string                    `json:"netBiosName,omitempty"`
	NetworkInterfaces         *[]NetworkInterface        `json:"networkInterfaces,omitempty"`
	ODataType                 *string                    `json:"@odata.type,omitempty"`
	Os                        *string                    `json:"os,omitempty"`
	OsVersion                 *string                    `json:"osVersion,omitempty"`
	ParentHost                *string                    `json:"parentHost,omitempty"`
	RelatedHostIds            *[]string                  `json:"relatedHostIds,omitempty"`
	RiskScore                 *string                    `json:"riskScore,omitempty"`
	Tags                      *[]string                  `json:"tags,omitempty"`
	VendorInformation         *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}
