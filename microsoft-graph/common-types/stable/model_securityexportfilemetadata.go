package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportFileMetadata struct {
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	FileName    *string `json:"fileName,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Size        *int64  `json:"size,omitempty"`
}
