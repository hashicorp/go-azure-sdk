package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobTriggerTypeProperties struct {
	FolderPath     string                 `json:"folderPath"`
	LinkedService  LinkedServiceReference `json:"linkedService"`
	MaxConcurrency int64                  `json:"maxConcurrency"`
}
