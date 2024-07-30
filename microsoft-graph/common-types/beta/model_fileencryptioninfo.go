package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileEncryptionInfo struct {
	EncryptionKey        *string `json:"encryptionKey,omitempty"`
	FileDigest           *string `json:"fileDigest,omitempty"`
	FileDigestAlgorithm  *string `json:"fileDigestAlgorithm,omitempty"`
	InitializationVector *string `json:"initializationVector,omitempty"`
	Mac                  *string `json:"mac,omitempty"`
	MacKey               *string `json:"macKey,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	ProfileIdentifier    *string `json:"profileIdentifier,omitempty"`
}
