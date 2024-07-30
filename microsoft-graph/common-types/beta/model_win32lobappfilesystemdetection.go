package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemDetection struct {
	Check32BitOn64System *bool                                        `json:"check32BitOn64System,omitempty"`
	DetectionType        *Win32LobAppFileSystemDetectionDetectionType `json:"detectionType,omitempty"`
	DetectionValue       *string                                      `json:"detectionValue,omitempty"`
	FileOrFolderName     *string                                      `json:"fileOrFolderName,omitempty"`
	ODataType            *string                                      `json:"@odata.type,omitempty"`
	Operator             *Win32LobAppFileSystemDetectionOperator      `json:"operator,omitempty"`
	Path                 *string                                      `json:"path,omitempty"`
}
