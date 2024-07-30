package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItem struct {
	Activities           *[]ItemActivityOLD  `json:"activities,omitempty"`
	Analytics            *ItemAnalytics      `json:"analytics,omitempty"`
	Audio                *Audio              `json:"audio,omitempty"`
	Bundle               *Bundle             `json:"bundle,omitempty"`
	CTag                 *string             `json:"cTag,omitempty"`
	Children             *[]DriveItem        `json:"children,omitempty"`
	Content              *string             `json:"content,omitempty"`
	CreatedBy            *IdentitySet        `json:"createdBy,omitempty"`
	CreatedByUser        *User               `json:"createdByUser,omitempty"`
	CreatedDateTime      *string             `json:"createdDateTime,omitempty"`
	Deleted              *Deleted            `json:"deleted,omitempty"`
	Description          *string             `json:"description,omitempty"`
	ETag                 *string             `json:"eTag,omitempty"`
	File                 *File               `json:"file,omitempty"`
	FileSystemInfo       *FileSystemInfo     `json:"fileSystemInfo,omitempty"`
	Folder               *Folder             `json:"folder,omitempty"`
	Id                   *string             `json:"id,omitempty"`
	Image                *Image              `json:"image,omitempty"`
	LastModifiedBy       *IdentitySet        `json:"lastModifiedBy,omitempty"`
	LastModifiedByUser   *User               `json:"lastModifiedByUser,omitempty"`
	LastModifiedDateTime *string             `json:"lastModifiedDateTime,omitempty"`
	ListItem             *ListItem           `json:"listItem,omitempty"`
	Location             *GeoCoordinates     `json:"location,omitempty"`
	Malware              *Malware            `json:"malware,omitempty"`
	Media                *Media              `json:"media,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	ODataType            *string             `json:"@odata.type,omitempty"`
	Package              *Package            `json:"package,omitempty"`
	ParentReference      *ItemReference      `json:"parentReference,omitempty"`
	PendingOperations    *PendingOperations  `json:"pendingOperations,omitempty"`
	Permissions          *[]Permission       `json:"permissions,omitempty"`
	Photo                *Photo              `json:"photo,omitempty"`
	Publication          *PublicationFacet   `json:"publication,omitempty"`
	RemoteItem           *RemoteItem         `json:"remoteItem,omitempty"`
	RetentionLabel       *ItemRetentionLabel `json:"retentionLabel,omitempty"`
	Root                 *Root               `json:"root,omitempty"`
	SearchResult         *SearchResult       `json:"searchResult,omitempty"`
	Shared               *Shared             `json:"shared,omitempty"`
	SharepointIds        *SharepointIds      `json:"sharepointIds,omitempty"`
	Size                 *int64              `json:"size,omitempty"`
	Source               *DriveItemSource    `json:"source,omitempty"`
	SpecialFolder        *SpecialFolder      `json:"specialFolder,omitempty"`
	Subscriptions        *[]Subscription     `json:"subscriptions,omitempty"`
	Thumbnails           *[]ThumbnailSet     `json:"thumbnails,omitempty"`
	Versions             *[]DriveItemVersion `json:"versions,omitempty"`
	Video                *Video              `json:"video,omitempty"`
	WebDavUrl            *string             `json:"webDavUrl,omitempty"`
	WebUrl               *string             `json:"webUrl,omitempty"`
	Workbook             *Workbook           `json:"workbook,omitempty"`
}
