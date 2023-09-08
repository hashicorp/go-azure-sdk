package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseSitePage struct {
	CreatedBy            *IdentitySet            `json:"createdBy,omitempty"`
	CreatedByUser        *User                   `json:"createdByUser,omitempty"`
	CreatedDateTime      *string                 `json:"createdDateTime,omitempty"`
	Description          *string                 `json:"description,omitempty"`
	ETag                 *string                 `json:"eTag,omitempty"`
	Id                   *string                 `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet            `json:"lastModifiedBy,omitempty"`
	LastModifiedByUser   *User                   `json:"lastModifiedByUser,omitempty"`
	LastModifiedDateTime *string                 `json:"lastModifiedDateTime,omitempty"`
	Name                 *string                 `json:"name,omitempty"`
	ODataType            *string                 `json:"@odata.type,omitempty"`
	PageLayout           *BaseSitePagePageLayout `json:"pageLayout,omitempty"`
	ParentReference      *ItemReference          `json:"parentReference,omitempty"`
	PublishingState      *PublicationFacet       `json:"publishingState,omitempty"`
	Title                *string                 `json:"title,omitempty"`
	WebUrl               *string                 `json:"webUrl,omitempty"`
}
