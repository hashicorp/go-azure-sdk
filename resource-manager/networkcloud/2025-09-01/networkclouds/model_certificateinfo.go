package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateInfo struct {
	Hash  *string `json:"hash,omitempty"`
	Value *string `json:"value,omitempty"`
}
