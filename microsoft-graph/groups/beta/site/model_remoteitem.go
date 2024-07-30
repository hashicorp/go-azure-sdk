package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteItem struct {
	CreatedBy            *IdentitySet    `json:"createdBy,omitempty"`
	CreatedDateTime      *string         `json:"createdDateTime,omitempty"`
	File                 *File           `json:"file,omitempty"`
	FileSystemInfo       *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	Folder               *Folder         `json:"folder,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	Image                *Image          `json:"image,omitempty"`
	LastModifiedBy       *IdentitySet    `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string         `json:"lastModifiedDateTime,omitempty"`
	Name                 *string         `json:"name,omitempty"`
	ODataType            *string         `json:"@odata.type,omitempty"`
	Package              *Package        `json:"package,omitempty"`
	ParentReference      *ItemReference  `json:"parentReference,omitempty"`
	Shared               *Shared         `json:"shared,omitempty"`
	SharepointIds        *SharepointIds  `json:"sharepointIds,omitempty"`
	Size                 *int64          `json:"size,omitempty"`
	SpecialFolder        *SpecialFolder  `json:"specialFolder,omitempty"`
	Video                *Video          `json:"video,omitempty"`
	WebDavUrl            *string         `json:"webDavUrl,omitempty"`
	WebUrl               *string         `json:"webUrl,omitempty"`
}
