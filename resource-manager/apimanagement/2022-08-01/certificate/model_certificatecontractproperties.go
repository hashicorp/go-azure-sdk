package certificate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateContractProperties struct {
	ExpirationDate string                      `json:"expirationDate"`
	KeyVault       *KeyVaultContractProperties `json:"keyVault,omitempty"`
	Subject        string                      `json:"subject"`
	Thumbprint     string                      `json:"thumbprint"`
}
