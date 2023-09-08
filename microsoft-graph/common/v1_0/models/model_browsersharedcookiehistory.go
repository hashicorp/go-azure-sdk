package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieHistory struct {
	Comment           *string                                      `json:"comment,omitempty"`
	DisplayName       *string                                      `json:"displayName,omitempty"`
	HostOnly          *bool                                        `json:"hostOnly,omitempty"`
	HostOrDomain      *string                                      `json:"hostOrDomain,omitempty"`
	LastModifiedBy    *IdentitySet                                 `json:"lastModifiedBy,omitempty"`
	ODataType         *string                                      `json:"@odata.type,omitempty"`
	Path              *string                                      `json:"path,omitempty"`
	PublishedDateTime *string                                      `json:"publishedDateTime,omitempty"`
	SourceEnvironment *BrowserSharedCookieHistorySourceEnvironment `json:"sourceEnvironment,omitempty"`
}
