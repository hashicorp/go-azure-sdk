package certificates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAddParameter struct {
	CertificateFormat   *CertificateFormat `json:"certificateFormat,omitempty"`
	Data                string             `json:"data"`
	Password            *string            `json:"password,omitempty"`
	Thumbprint          string             `json:"thumbprint"`
	ThumbprintAlgorithm string             `json:"thumbprintAlgorithm"`
}
