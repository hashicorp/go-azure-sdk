package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SigningX509Certificate struct {
	Certificate *string              `json:"certificate,omitempty"`
	Etag        *string              `json:"etag,omitempty"`
	Info        *X509CertificateInfo `json:"info,omitempty"`
	Verified    *bool                `json:"verified,omitempty"`
}
