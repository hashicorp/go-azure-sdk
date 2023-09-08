package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Pkcs12CertificateInformation struct {
	IsActive   *bool   `json:"isActive,omitempty"`
	NotAfter   *int64  `json:"notAfter,omitempty"`
	NotBefore  *int64  `json:"notBefore,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	Thumbprint *string `json:"thumbprint,omitempty"`
}
