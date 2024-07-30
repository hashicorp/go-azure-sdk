package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsTabConfiguration struct {
	ContentUrl *string `json:"contentUrl,omitempty"`
	EntityId   *string `json:"entityId,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	RemoveUrl  *string `json:"removeUrl,omitempty"`
	WebsiteUrl *string `json:"websiteUrl,omitempty"`
}
