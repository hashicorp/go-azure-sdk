package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderTenantSetting struct {
	AzureTenantId        *string `json:"azureTenantId,omitempty"`
	Enabled              *bool   `json:"enabled,omitempty"`
	Id                   *string `json:"id,omitempty"`
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	Provider             *string `json:"provider,omitempty"`
	Vendor               *string `json:"vendor,omitempty"`
}
