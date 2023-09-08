package groupsiteinformationprotection

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupByIdSiteByIdInformationProtectionVerifySignatureRequest struct {
	Digest       *string `json:"digest,omitempty"`
	Signature    *string `json:"signature,omitempty"`
	SigningKeyId *string `json:"signingKeyId,omitempty"`
}
