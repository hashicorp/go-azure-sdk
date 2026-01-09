package informationprotection

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionDecryptBufferRequest struct {
	EncryptedBuffer   *string `json:"encryptedBuffer,omitempty"`
	PublishingLicense *string `json:"publishingLicense,omitempty"`
}
