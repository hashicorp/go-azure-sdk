package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProviderStatus struct {
	Enabled   *bool   `json:"enabled,omitempty"`
	Endpoint  *string `json:"endpoint,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Provider  *string `json:"provider,omitempty"`
	Region    *string `json:"region,omitempty"`
	Vendor    *string `json:"vendor,omitempty"`
}
