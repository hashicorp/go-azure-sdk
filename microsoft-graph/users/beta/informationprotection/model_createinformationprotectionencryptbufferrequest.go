package informationprotection

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionEncryptBufferRequest struct {
	Buffer  *string `json:"buffer,omitempty"`
	LabelId *string `json:"labelId,omitempty"`
}
