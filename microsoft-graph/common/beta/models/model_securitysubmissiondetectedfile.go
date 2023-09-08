package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionDetectedFile struct {
	FileHash  *string `json:"fileHash,omitempty"`
	FileName  *string `json:"fileName,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
