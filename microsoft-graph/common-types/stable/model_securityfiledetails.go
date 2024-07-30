package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileDetails struct {
	FileName      *string `json:"fileName,omitempty"`
	FilePath      *string `json:"filePath,omitempty"`
	FilePublisher *string `json:"filePublisher,omitempty"`
	FileSize      *int64  `json:"fileSize,omitempty"`
	Issuer        *string `json:"issuer,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	Sha1          *string `json:"sha1,omitempty"`
	Sha256        *string `json:"sha256,omitempty"`
	Signer        *string `json:"signer,omitempty"`
}
