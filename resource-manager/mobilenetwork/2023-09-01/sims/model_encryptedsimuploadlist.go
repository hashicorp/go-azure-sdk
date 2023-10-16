package sims

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptedSimUploadList struct {
	AzureKeyIdentifier    int64                           `json:"azureKeyIdentifier"`
	EncryptedTransportKey string                          `json:"encryptedTransportKey"`
	SignedTransportKey    string                          `json:"signedTransportKey"`
	Sims                  []SimNameAndEncryptedProperties `json:"sims"`
	VendorKeyFingerprint  string                          `json:"vendorKeyFingerprint"`
	Version               int64                           `json:"version"`
}
