package fileimports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileMetadata struct {
	DeleteStatus   *DeleteStatus `json:"deleteStatus,omitempty"`
	FileContentUri *string       `json:"fileContentUri,omitempty"`
	FileFormat     *FileFormat   `json:"fileFormat,omitempty"`
	FileName       *string       `json:"fileName,omitempty"`
	FileSize       *int64        `json:"fileSize,omitempty"`
}
