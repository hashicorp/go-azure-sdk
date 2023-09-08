package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Drive struct {
	Bundles              *[]DriveItem   `json:"bundles,omitempty"`
	CreatedBy            *IdentitySet   `json:"createdBy,omitempty"`
	CreatedByUser        *User          `json:"createdByUser,omitempty"`
	CreatedDateTime      *string        `json:"createdDateTime,omitempty"`
	Description          *string        `json:"description,omitempty"`
	DriveType            *string        `json:"driveType,omitempty"`
	ETag                 *string        `json:"eTag,omitempty"`
	Following            *[]DriveItem   `json:"following,omitempty"`
	Id                   *string        `json:"id,omitempty"`
	Items                *[]DriveItem   `json:"items,omitempty"`
	LastModifiedBy       *IdentitySet   `json:"lastModifiedBy,omitempty"`
	LastModifiedByUser   *User          `json:"lastModifiedByUser,omitempty"`
	LastModifiedDateTime *string        `json:"lastModifiedDateTime,omitempty"`
	List                 *List          `json:"list,omitempty"`
	Name                 *string        `json:"name,omitempty"`
	ODataType            *string        `json:"@odata.type,omitempty"`
	Owner                *IdentitySet   `json:"owner,omitempty"`
	ParentReference      *ItemReference `json:"parentReference,omitempty"`
	Quota                *Quota         `json:"quota,omitempty"`
	Root                 *DriveItem     `json:"root,omitempty"`
	SharePointIds        *SharepointIds `json:"sharePointIds,omitempty"`
	Special              *[]DriveItem   `json:"special,omitempty"`
	System               *SystemFacet   `json:"system,omitempty"`
	WebUrl               *string        `json:"webUrl,omitempty"`
}
