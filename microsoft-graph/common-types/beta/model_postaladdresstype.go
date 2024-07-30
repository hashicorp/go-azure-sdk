package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostalAddressType struct {
	City              *string `json:"city,omitempty"`
	CountryLetterCode *string `json:"countryLetterCode,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	PostalCode        *string `json:"postalCode,omitempty"`
	State             *string `json:"state,omitempty"`
	Street            *string `json:"street,omitempty"`
}
