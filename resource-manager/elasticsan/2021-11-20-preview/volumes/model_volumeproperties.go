package volumes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumeProperties struct {
	CreationData  *SourceCreationData `json:"creationData"`
	SizeGiB       *int64              `json:"sizeGiB,omitempty"`
	StorageTarget *IscsiTargetInfo    `json:"storageTarget"`
	VolumeId      *string             `json:"volumeId,omitempty"`
}
