package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSite struct {
	AllowRedirect        *bool                         `json:"allowRedirect,omitempty"`
	Comment              *string                       `json:"comment,omitempty"`
	CompatibilityMode    *BrowserSiteCompatibilityMode `json:"compatibilityMode,omitempty"`
	CreatedDateTime      *string                       `json:"createdDateTime,omitempty"`
	DeletedDateTime      *string                       `json:"deletedDateTime,omitempty"`
	History              *[]BrowserSiteHistory         `json:"history,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet                  `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                       `json:"lastModifiedDateTime,omitempty"`
	MergeType            *BrowserSiteMergeType         `json:"mergeType,omitempty"`
	ODataType            *string                       `json:"@odata.type,omitempty"`
	Status               *BrowserSiteStatus            `json:"status,omitempty"`
	TargetEnvironment    *BrowserSiteTargetEnvironment `json:"targetEnvironment,omitempty"`
	WebUrl               *string                       `json:"webUrl,omitempty"`
}
