package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Site struct {
	Analytics            *ItemAnalytics              `json:"analytics,omitempty"`
	Columns              *[]ColumnDefinition         `json:"columns,omitempty"`
	ContentTypes         *[]ContentType              `json:"contentTypes,omitempty"`
	CreatedBy            *IdentitySet                `json:"createdBy,omitempty"`
	CreatedByUser        *User                       `json:"createdByUser,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	DisplayName          *string                     `json:"displayName,omitempty"`
	Drive                *Drive                      `json:"drive,omitempty"`
	Drives               *[]Drive                    `json:"drives,omitempty"`
	ETag                 *string                     `json:"eTag,omitempty"`
	Error                *PublicError                `json:"error,omitempty"`
	ExternalColumns      *[]ColumnDefinition         `json:"externalColumns,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	IsPersonalSite       *bool                       `json:"isPersonalSite,omitempty"`
	Items                *[]BaseItem                 `json:"items,omitempty"`
	LastModifiedBy       *IdentitySet                `json:"lastModifiedBy,omitempty"`
	LastModifiedByUser   *User                       `json:"lastModifiedByUser,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	Lists                *[]List                     `json:"lists,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Onenote              *Onenote                    `json:"onenote,omitempty"`
	Operations           *[]RichLongRunningOperation `json:"operations,omitempty"`
	ParentReference      *ItemReference              `json:"parentReference,omitempty"`
	Permissions          *[]Permission               `json:"permissions,omitempty"`
	Root                 *Root                       `json:"root,omitempty"`
	SharepointIds        *SharepointIds              `json:"sharepointIds,omitempty"`
	SiteCollection       *SiteCollection             `json:"siteCollection,omitempty"`
	Sites                *[]Site                     `json:"sites,omitempty"`
	TermStore            *TermStoreStore             `json:"termStore,omitempty"`
	TermStores           *[]TermStoreStore           `json:"termStores,omitempty"`
	WebUrl               *string                     `json:"webUrl,omitempty"`
}
