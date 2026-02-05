package files

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeFile struct {
	IsDirectory *bool           `json:"isDirectory,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Properties  *FileProperties `json:"properties,omitempty"`
	Url         *string         `json:"url,omitempty"`
}
