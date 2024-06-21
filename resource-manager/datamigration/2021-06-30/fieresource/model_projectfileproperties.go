package fieresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectFileProperties struct {
	Extension    *string `json:"extension,omitempty"`
	FilePath     *string `json:"filePath,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	MediaType    *string `json:"mediaType,omitempty"`
	Size         *int64  `json:"size,omitempty"`
}
