package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultCertificate struct {
	CertificateStore *string `json:"certificateStore,omitempty"`
	CertificateURL   *string `json:"certificateUrl,omitempty"`
}
