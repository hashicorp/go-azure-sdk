package galleryapplicationversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GalleryApplicationVersionPublishingProfile struct {
	EnableHealthCheck  *bool               `json:"enableHealthCheck,omitempty"`
	EndOfLifeDate      *string             `json:"endOfLifeDate,omitempty"`
	ExcludeFromLatest  *bool               `json:"excludeFromLatest,omitempty"`
	ManageActions      *UserArtifactManage `json:"manageActions,omitempty"`
	PublishedDate      *string             `json:"publishedDate,omitempty"`
	ReplicaCount       *int64              `json:"replicaCount,omitempty"`
	ReplicationMode    *ReplicationMode    `json:"replicationMode,omitempty"`
	Source             UserArtifactSource  `json:"source"`
	StorageAccountType *StorageAccountType `json:"storageAccountType,omitempty"`
	TargetRegions      *[]TargetRegion     `json:"targetRegions,omitempty"`
}
