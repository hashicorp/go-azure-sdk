package sharedgalleryimageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleryDataDiskImage struct {
	DiskSizeGB  *int64                    `json:"diskSizeGB,omitempty"`
	HostCaching *SharedGalleryHostCaching `json:"hostCaching,omitempty"`
	Lun         int64                     `json:"lun"`
}
