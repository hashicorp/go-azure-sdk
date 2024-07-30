package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySslCertificateEntity struct {
	Address              *PhysicalAddress `json:"address,omitempty"`
	AlternateNames       *[]string        `json:"alternateNames,omitempty"`
	CommonName           *string          `json:"commonName,omitempty"`
	Email                *string          `json:"email,omitempty"`
	GivenName            *string          `json:"givenName,omitempty"`
	ODataType            *string          `json:"@odata.type,omitempty"`
	OrganizationName     *string          `json:"organizationName,omitempty"`
	OrganizationUnitName *string          `json:"organizationUnitName,omitempty"`
	SerialNumber         *string          `json:"serialNumber,omitempty"`
	Surname              *string          `json:"surname,omitempty"`
}
