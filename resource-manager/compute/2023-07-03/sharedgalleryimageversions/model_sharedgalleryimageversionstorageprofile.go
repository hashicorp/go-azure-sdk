package sharedgalleryimageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleryImageVersionStorageProfile struct {
	DataDiskImages *[]SharedGalleryDataDiskImage `json:"dataDiskImages,omitempty"`
	OsDiskImage    *SharedGalleryDiskImage       `json:"osDiskImage,omitempty"`
}
