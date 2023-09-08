package groupsitecontenttype

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupByIdSiteByIdContentTypeByIdAssociateWithHubSiteRequest struct {
	HubSiteUrls              *[]string `json:"hubSiteUrls,omitempty"`
	PropagateToExistingLists *bool     `json:"propagateToExistingLists,omitempty"`
}
