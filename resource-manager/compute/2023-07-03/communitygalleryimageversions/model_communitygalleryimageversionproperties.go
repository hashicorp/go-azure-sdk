package communitygalleryimageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageVersionProperties struct {
	ArtifactTags      *map[string]string                       `json:"artifactTags,omitempty"`
	Disclaimer        *CommunityGalleryDisclaimer              `json:"disclaimer,omitempty"`
	EndOfLifeDate     *string                                  `json:"endOfLifeDate,omitempty"`
	ExcludeFromLatest *bool                                    `json:"excludeFromLatest,omitempty"`
	PublishedDate     *string                                  `json:"publishedDate,omitempty"`
	StorageProfile    *SharedGalleryImageVersionStorageProfile `json:"storageProfile,omitempty"`
}
