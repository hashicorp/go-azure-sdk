package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityProfile struct {
	Accounts             *[]UserAccount             `json:"accounts,omitempty"`
	AzureSubscriptionId  *string                    `json:"azureSubscriptionId,omitempty"`
	AzureTenantId        *string                    `json:"azureTenantId,omitempty"`
	CreatedDateTime      *string                    `json:"createdDateTime,omitempty"`
	DisplayName          *string                    `json:"displayName,omitempty"`
	Id                   *string                    `json:"id,omitempty"`
	LastModifiedDateTime *string                    `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                    `json:"@odata.type,omitempty"`
	RiskScore            *string                    `json:"riskScore,omitempty"`
	Tags                 *[]string                  `json:"tags,omitempty"`
	UserPrincipalName    *string                    `json:"userPrincipalName,omitempty"`
	VendorInformation    *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}
