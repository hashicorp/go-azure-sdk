package informationprotection

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionEncryptBufferRequest struct {
	Buffer  *string `json:"buffer,omitempty"`
	LabelId *string `json:"labelId,omitempty"`
}
