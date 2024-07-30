package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookie struct {
	Comment              *string                               `json:"comment,omitempty"`
	CreatedDateTime      *string                               `json:"createdDateTime,omitempty"`
	DeletedDateTime      *string                               `json:"deletedDateTime,omitempty"`
	DisplayName          *string                               `json:"displayName,omitempty"`
	History              *[]BrowserSharedCookieHistory         `json:"history,omitempty"`
	HostOnly             *bool                                 `json:"hostOnly,omitempty"`
	HostOrDomain         *string                               `json:"hostOrDomain,omitempty"`
	Id                   *string                               `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet                          `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                               `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	Path                 *string                               `json:"path,omitempty"`
	SourceEnvironment    *BrowserSharedCookieSourceEnvironment `json:"sourceEnvironment,omitempty"`
	Status               *BrowserSharedCookieStatus            `json:"status,omitempty"`
}
