package apimanagementservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateInformation struct {
	Expiry     string `json:"expiry"`
	Subject    string `json:"subject"`
	Thumbprint string `json:"thumbprint"`
}
