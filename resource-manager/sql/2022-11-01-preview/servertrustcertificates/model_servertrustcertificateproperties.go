package servertrustcertificates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerTrustCertificateProperties struct {
	CertificateName *string `json:"certificateName,omitempty"`
	PublicBlob      *string `json:"publicBlob,omitempty"`
	Thumbprint      *string `json:"thumbprint,omitempty"`
}
