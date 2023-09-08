package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePage struct {
	CanvasLayout         *CanvasLayout          `json:"canvasLayout,omitempty"`
	CreatedBy            *IdentitySet           `json:"createdBy,omitempty"`
	CreatedByUser        *User                  `json:"createdByUser,omitempty"`
	CreatedDateTime      *string                `json:"createdDateTime,omitempty"`
	Description          *string                `json:"description,omitempty"`
	ETag                 *string                `json:"eTag,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet           `json:"lastModifiedBy,omitempty"`
	LastModifiedByUser   *User                  `json:"lastModifiedByUser,omitempty"`
	LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	PageLayout           *SitePagePageLayout    `json:"pageLayout,omitempty"`
	ParentReference      *ItemReference         `json:"parentReference,omitempty"`
	PromotionKind        *SitePagePromotionKind `json:"promotionKind,omitempty"`
	PublishingState      *PublicationFacet      `json:"publishingState,omitempty"`
	Reactions            *ReactionsFacet        `json:"reactions,omitempty"`
	ShowComments         *bool                  `json:"showComments,omitempty"`
	ShowRecommendedPages *bool                  `json:"showRecommendedPages,omitempty"`
	ThumbnailWebUrl      *string                `json:"thumbnailWebUrl,omitempty"`
	Title                *string                `json:"title,omitempty"`
	TitleArea            *TitleArea             `json:"titleArea,omitempty"`
	WebParts             *[]WebPart             `json:"webParts,omitempty"`
	WebUrl               *string                `json:"webUrl,omitempty"`
}
