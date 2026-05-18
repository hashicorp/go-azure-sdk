package siteinformationprotection

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteInformationProtectionVerifySignatureRequest struct {
	Digest       *string `json:"digest,omitempty"`
	Signature    *string `json:"signature,omitempty"`
	SigningKeyId *string `json:"signingKeyId,omitempty"`
}
