package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateCredentialsUpdate struct {
	CertificateSecretName              *string `json:"certificateSecretName,omitempty"`
	IntermediateCertificatesSecretName *string `json:"intermediateCertificatesSecretName,omitempty"`
	KeySecretName                      *string `json:"keySecretName,omitempty"`
}
