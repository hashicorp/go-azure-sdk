package galleryimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GalleryImageProperties struct {
	Author           *string                `json:"author,omitempty"`
	CreatedDate      *string                `json:"createdDate,omitempty"`
	Description      *string                `json:"description,omitempty"`
	Enabled          *bool                  `json:"enabled,omitempty"`
	Icon             *string                `json:"icon,omitempty"`
	ImageReference   *GalleryImageReference `json:"imageReference,omitempty"`
	IsPlanAuthorized *bool                  `json:"isPlanAuthorized,omitempty"`
	PlanId           *string                `json:"planId,omitempty"`
}
