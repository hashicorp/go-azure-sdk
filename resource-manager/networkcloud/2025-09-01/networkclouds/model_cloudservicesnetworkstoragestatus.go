package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudServicesNetworkStorageStatus struct {
	Mode          *CloudServicesNetworkStorageMode         `json:"mode,omitempty"`
	SizeMiB       *int64                                   `json:"sizeMiB,omitempty"`
	Status        *CloudServicesNetworkStorageStatusStatus `json:"status,omitempty"`
	StatusMessage *string                                  `json:"statusMessage,omitempty"`
	VolumeId      *string                                  `json:"volumeId,omitempty"`
}
