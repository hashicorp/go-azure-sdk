package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteList struct {
	Description          *string                `json:"description,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet           `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	PublishedBy          *IdentitySet           `json:"publishedBy,omitempty"`
	PublishedDateTime    *string                `json:"publishedDateTime,omitempty"`
	Revision             *string                `json:"revision,omitempty"`
	SharedCookies        *[]BrowserSharedCookie `json:"sharedCookies,omitempty"`
	Sites                *[]BrowserSite         `json:"sites,omitempty"`
	Status               *BrowserSiteListStatus `json:"status,omitempty"`
}
