package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharePointIdentitySet struct {
	Application *Identity           `json:"application,omitempty"`
	Device      *Identity           `json:"device,omitempty"`
	Group       *Identity           `json:"group,omitempty"`
	ODataType   *string             `json:"@odata.type,omitempty"`
	SiteGroup   *SharePointIdentity `json:"siteGroup,omitempty"`
	SiteUser    *SharePointIdentity `json:"siteUser,omitempty"`
	User        *Identity           `json:"user,omitempty"`
}
