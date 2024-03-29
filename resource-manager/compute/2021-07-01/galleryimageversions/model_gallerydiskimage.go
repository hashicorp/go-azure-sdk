package galleryimageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GalleryDiskImage struct {
	HostCaching *HostCaching                  `json:"hostCaching,omitempty"`
	SizeInGB    *int64                        `json:"sizeInGB,omitempty"`
	Source      *GalleryArtifactVersionSource `json:"source,omitempty"`
}
