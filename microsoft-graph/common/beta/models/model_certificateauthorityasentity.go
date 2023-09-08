package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityAsEntity struct {
	Certificate                *string `json:"certificate,omitempty"`
	Id                         *string `json:"id,omitempty"`
	IsRootAuthority            *bool   `json:"isRootAuthority,omitempty"`
	Issuer                     *string `json:"issuer,omitempty"`
	IssuerSubjectKeyIdentifier *string `json:"issuerSubjectKeyIdentifier,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
}
