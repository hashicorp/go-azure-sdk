package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type List struct {
	Columns              *[]ColumnDefinition         `json:"columns,omitempty"`
	ContentTypes         *[]ContentType              `json:"contentTypes,omitempty"`
	CreatedBy            *IdentitySet                `json:"createdBy,omitempty"`
	CreatedByUser        *User                       `json:"createdByUser,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	DisplayName          *string                     `json:"displayName,omitempty"`
	Drive                *Drive                      `json:"drive,omitempty"`
	ETag                 *string                     `json:"eTag,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Items                *[]ListItem                 `json:"items,omitempty"`
	LastModifiedBy       *IdentitySet                `json:"lastModifiedBy,omitempty"`
	LastModifiedByUser   *User                       `json:"lastModifiedByUser,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	List                 *ListInfo                   `json:"list,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Operations           *[]RichLongRunningOperation `json:"operations,omitempty"`
	ParentReference      *ItemReference              `json:"parentReference,omitempty"`
	SharepointIds        *SharepointIds              `json:"sharepointIds,omitempty"`
	Subscriptions        *[]Subscription             `json:"subscriptions,omitempty"`
	System               *SystemFacet                `json:"system,omitempty"`
	WebUrl               *string                     `json:"webUrl,omitempty"`
}
