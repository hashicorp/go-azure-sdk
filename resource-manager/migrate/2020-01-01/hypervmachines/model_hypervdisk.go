package hypervmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVDisk struct {
	DiskType       *string `json:"diskType,omitempty"`
	InstanceId     *string `json:"instanceId,omitempty"`
	Lun            *int64  `json:"lun,omitempty"`
	MaxSizeInBytes *int64  `json:"maxSizeInBytes,omitempty"`
	Name           *string `json:"name,omitempty"`
	Path           *string `json:"path,omitempty"`
	VhdId          *string `json:"vhdId,omitempty"`
}
