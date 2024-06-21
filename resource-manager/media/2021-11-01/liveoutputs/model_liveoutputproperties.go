package liveoutputs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiveOutputProperties struct {
	ArchiveWindowLength string                   `json:"archiveWindowLength"`
	AssetName           string                   `json:"assetName"`
	Created             *string                  `json:"created,omitempty"`
	Description         *string                  `json:"description,omitempty"`
	Hls                 *Hls                     `json:"hls,omitempty"`
	LastModified        *string                  `json:"lastModified,omitempty"`
	ManifestName        *string                  `json:"manifestName,omitempty"`
	OutputSnapTime      *int64                   `json:"outputSnapTime,omitempty"`
	ProvisioningState   *string                  `json:"provisioningState,omitempty"`
	ResourceState       *LiveOutputResourceState `json:"resourceState,omitempty"`
}
