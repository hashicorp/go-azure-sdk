package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointIds struct {
	ListId           *string `json:"listId,omitempty"`
	ListItemId       *string `json:"listItemId,omitempty"`
	ListItemUniqueId *string `json:"listItemUniqueId,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	SiteId           *string `json:"siteId,omitempty"`
	SiteUrl          *string `json:"siteUrl,omitempty"`
	TenantId         *string `json:"tenantId,omitempty"`
	WebId            *string `json:"webId,omitempty"`
}
