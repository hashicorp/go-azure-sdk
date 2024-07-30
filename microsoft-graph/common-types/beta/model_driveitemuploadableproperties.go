package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemUploadableProperties struct {
	Description    *string         `json:"description,omitempty"`
	FileSize       *int64          `json:"fileSize,omitempty"`
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	Name           *string         `json:"name,omitempty"`
	ODataType      *string         `json:"@odata.type,omitempty"`
}
