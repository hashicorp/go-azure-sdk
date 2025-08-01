package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageProfile struct {
	AlignRegionalDisksToVMZone *bool                `json:"alignRegionalDisksToVMZone,omitempty"`
	DataDisks                  *[]DataDisk          `json:"dataDisks,omitempty"`
	DiskControllerType         *DiskControllerTypes `json:"diskControllerType,omitempty"`
	ImageReference             *ImageReference      `json:"imageReference,omitempty"`
	OsDisk                     *OSDisk              `json:"osDisk,omitempty"`
}
