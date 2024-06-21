package tokens

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenCertificate struct {
	EncodedPemCertificate *string               `json:"encodedPemCertificate,omitempty"`
	Expiry                *string               `json:"expiry,omitempty"`
	Name                  *TokenCertificateName `json:"name,omitempty"`
	Thumbprint            *string               `json:"thumbprint,omitempty"`
}
