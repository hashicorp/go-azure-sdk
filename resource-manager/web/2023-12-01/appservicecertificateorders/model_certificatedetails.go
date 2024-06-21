package appservicecertificateorders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateDetails struct {
	Issuer             *string `json:"issuer,omitempty"`
	NotAfter           *string `json:"notAfter,omitempty"`
	NotBefore          *string `json:"notBefore,omitempty"`
	RawData            *string `json:"rawData,omitempty"`
	SerialNumber       *string `json:"serialNumber,omitempty"`
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	Subject            *string `json:"subject,omitempty"`
	Thumbprint         *string `json:"thumbprint,omitempty"`
	Version            *int64  `json:"version,omitempty"`
}
