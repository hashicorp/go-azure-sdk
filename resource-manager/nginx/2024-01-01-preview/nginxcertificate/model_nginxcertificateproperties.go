package nginxcertificate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NginxCertificateProperties struct {
	CertificateError       *NginxCertificateErrorResponseBody `json:"certificateError,omitempty"`
	CertificateVirtualPath *string                            `json:"certificateVirtualPath,omitempty"`
	KeyVaultSecretCreated  *string                            `json:"keyVaultSecretCreated,omitempty"`
	KeyVaultSecretId       *string                            `json:"keyVaultSecretId,omitempty"`
	KeyVaultSecretVersion  *string                            `json:"keyVaultSecretVersion,omitempty"`
	KeyVirtualPath         *string                            `json:"keyVirtualPath,omitempty"`
	ProvisioningState      *ProvisioningState                 `json:"provisioningState,omitempty"`
	Sha1Thumbprint         *string                            `json:"sha1Thumbprint,omitempty"`
}
