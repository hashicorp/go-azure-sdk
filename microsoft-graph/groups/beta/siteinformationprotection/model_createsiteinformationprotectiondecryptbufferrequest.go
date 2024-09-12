package siteinformationprotection

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteInformationProtectionDecryptBufferRequest struct {
	EncryptedBuffer   *string `json:"encryptedBuffer,omitempty"`
	PublishingLicense *string `json:"publishingLicense,omitempty"`
}
