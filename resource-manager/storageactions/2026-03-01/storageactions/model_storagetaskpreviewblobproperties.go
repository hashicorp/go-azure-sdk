package storageactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageTaskPreviewBlobProperties struct {
	MatchedBlock *MatchedBlockName                       `json:"matchedBlock,omitempty"`
	Metadata     *[]StorageTaskPreviewKeyValueProperties `json:"metadata,omitempty"`
	Name         *string                                 `json:"name,omitempty"`
	Properties   *[]StorageTaskPreviewKeyValueProperties `json:"properties,omitempty"`
	Tags         *[]StorageTaskPreviewKeyValueProperties `json:"tags,omitempty"`
}
