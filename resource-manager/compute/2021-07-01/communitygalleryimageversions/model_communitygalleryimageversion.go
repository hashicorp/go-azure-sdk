package communitygalleryimageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageVersion struct {
	Identifier *CommunityGalleryIdentifier             `json:"identifier"`
	Location   *string                                 `json:"location,omitempty"`
	Name       *string                                 `json:"name,omitempty"`
	Properties *CommunityGalleryImageVersionProperties `json:"properties"`
	Type       *string                                 `json:"type,omitempty"`
}
