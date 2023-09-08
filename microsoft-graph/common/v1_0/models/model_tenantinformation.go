package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantInformation struct {
	DefaultDomainName   *string `json:"defaultDomainName,omitempty"`
	DisplayName         *string `json:"displayName,omitempty"`
	FederationBrandName *string `json:"federationBrandName,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	TenantId            *string `json:"tenantId,omitempty"`
}
