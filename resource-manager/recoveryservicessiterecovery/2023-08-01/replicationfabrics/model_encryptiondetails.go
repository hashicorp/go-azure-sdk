package replicationfabrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionDetails struct {
	KekCertExpiryDate *string `json:"kekCertExpiryDate,omitempty"`
	KekCertThumbprint *string `json:"kekCertThumbprint,omitempty"`
	KekState          *string `json:"kekState,omitempty"`
}
