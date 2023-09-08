package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpSecurityProfile struct {
	ActivityGroupNames  *[]string                  `json:"activityGroupNames,omitempty"`
	Address             *string                    `json:"address,omitempty"`
	AzureSubscriptionId *string                    `json:"azureSubscriptionId,omitempty"`
	AzureTenantId       *string                    `json:"azureTenantId,omitempty"`
	CountHits           *int64                     `json:"countHits,omitempty"`
	CountHosts          *int64                     `json:"countHosts,omitempty"`
	FirstSeenDateTime   *string                    `json:"firstSeenDateTime,omitempty"`
	Id                  *string                    `json:"id,omitempty"`
	IpCategories        *[]IpCategory              `json:"ipCategories,omitempty"`
	IpReferenceData     *[]IpReferenceData         `json:"ipReferenceData,omitempty"`
	LastSeenDateTime    *string                    `json:"lastSeenDateTime,omitempty"`
	ODataType           *string                    `json:"@odata.type,omitempty"`
	RiskScore           *string                    `json:"riskScore,omitempty"`
	Tags                *[]string                  `json:"tags,omitempty"`
	VendorInformation   *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}
