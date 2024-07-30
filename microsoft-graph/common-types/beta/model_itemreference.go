package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemReference struct {
	DriveId       *string        `json:"driveId,omitempty"`
	DriveType     *string        `json:"driveType,omitempty"`
	Id            *string        `json:"id,omitempty"`
	Name          *string        `json:"name,omitempty"`
	ODataType     *string        `json:"@odata.type,omitempty"`
	Path          *string        `json:"path,omitempty"`
	ShareId       *string        `json:"shareId,omitempty"`
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
	SiteId        *string        `json:"siteId,omitempty"`
}
