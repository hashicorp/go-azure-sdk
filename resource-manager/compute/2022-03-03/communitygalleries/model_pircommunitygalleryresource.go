package communitygalleries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PirCommunityGalleryResource struct {
	Identifier *CommunityGalleryIdentifier `json:"identifier,omitempty"`
	Location   *string                     `json:"location,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}
