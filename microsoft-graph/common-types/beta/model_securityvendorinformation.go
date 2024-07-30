package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVendorInformation struct {
	ODataType       *string `json:"@odata.type,omitempty"`
	Provider        *string `json:"provider,omitempty"`
	ProviderVersion *string `json:"providerVersion,omitempty"`
	SubProvider     *string `json:"subProvider,omitempty"`
	Vendor          *string `json:"vendor,omitempty"`
}
