package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthority struct {
	Certificate                       *string `json:"certificate,omitempty"`
	CertificateRevocationListUrl      *string `json:"certificateRevocationListUrl,omitempty"`
	DeltaCertificateRevocationListUrl *string `json:"deltaCertificateRevocationListUrl,omitempty"`
	IsRootAuthority                   *bool   `json:"isRootAuthority,omitempty"`
	Issuer                            *string `json:"issuer,omitempty"`
	IssuerSki                         *string `json:"issuerSki,omitempty"`
	ODataType                         *string `json:"@odata.type,omitempty"`
}
