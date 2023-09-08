package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BufferEncryptionResult struct {
	EncryptedBuffer   *string `json:"encryptedBuffer,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	PublishingLicense *string `json:"publishingLicense,omitempty"`
}
