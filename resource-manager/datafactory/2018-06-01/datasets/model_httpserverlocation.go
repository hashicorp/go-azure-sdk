package datasets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HTTPServerLocation struct {
	FileName    *interface{} `json:"fileName,omitempty"`
	FolderPath  *interface{} `json:"folderPath,omitempty"`
	RelativeUrl *interface{} `json:"relativeUrl,omitempty"`
	Type        string       `json:"type"`
}
