package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainRegistrant struct {
	CountryOrRegionCode *string `json:"countryOrRegionCode,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	Organization        *string `json:"organization,omitempty"`
	Url                 *string `json:"url,omitempty"`
	Vendor              *string `json:"vendor,omitempty"`
}
