package attestationproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationServiceCreationSpecificParams struct {
	PolicySigningCertificates    *JsonWebKeySet                    `json:"policySigningCertificates,omitempty"`
	PublicNetworkAccess          *PublicNetworkAccessType          `json:"publicNetworkAccess,omitempty"`
	TpmAttestationAuthentication *TpmAttestationAuthenticationType `json:"tpmAttestationAuthentication,omitempty"`
}
