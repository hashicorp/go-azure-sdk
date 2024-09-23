package informationprotection

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionSignDigestRequest struct {
	Digest *string `json:"digest,omitempty"`
}
