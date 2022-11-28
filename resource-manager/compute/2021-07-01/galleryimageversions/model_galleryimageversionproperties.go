package galleryimageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GalleryImageVersionProperties struct {
	ProvisioningState *ProvisioningState                    `json:"provisioningState,omitempty"`
	PublishingProfile *GalleryArtifactPublishingProfileBase `json:"publishingProfile"`
	ReplicationStatus *ReplicationStatus                    `json:"replicationStatus"`
	StorageProfile    GalleryImageVersionStorageProfile     `json:"storageProfile"`
}
