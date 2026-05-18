package siteinformationprotection

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteInformationProtectionDecryptBufferRequest struct {
	EncryptedBuffer   *string `json:"encryptedBuffer,omitempty"`
	PublishingLicense *string `json:"publishingLicense,omitempty"`
}
