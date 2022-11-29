package communitygalleryimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImage struct {
	Identifier *CommunityGalleryIdentifier      `json:"identifier,omitempty"`
	Location   *string                          `json:"location,omitempty"`
	Name       *string                          `json:"name,omitempty"`
	Properties *CommunityGalleryImageProperties `json:"properties,omitempty"`
	Type       *string                          `json:"type,omitempty"`
}
