package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpReferenceData struct {
	Asn                 *int64  `json:"asn,omitempty"`
	City                *string `json:"city,omitempty"`
	CountryOrRegionCode *string `json:"countryOrRegionCode,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	Organization        *string `json:"organization,omitempty"`
	State               *string `json:"state,omitempty"`
	Vendor              *string `json:"vendor,omitempty"`
}
