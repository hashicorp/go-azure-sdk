package storageactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageTaskPreviewContainerProperties struct {
	Metadata *[]StorageTaskPreviewKeyValueProperties `json:"metadata,omitempty"`
	Name     *string                                 `json:"name,omitempty"`
}
