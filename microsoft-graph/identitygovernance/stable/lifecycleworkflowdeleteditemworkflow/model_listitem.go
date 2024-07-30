package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListItem struct {
	Analytics            *ItemAnalytics        `json:"analytics,omitempty"`
	ContentType          *ContentTypeInfo      `json:"contentType,omitempty"`
	CreatedBy            *IdentitySet          `json:"createdBy,omitempty"`
	CreatedByUser        *User                 `json:"createdByUser,omitempty"`
	CreatedDateTime      *string               `json:"createdDateTime,omitempty"`
	Description          *string               `json:"description,omitempty"`
	DocumentSetVersions  *[]DocumentSetVersion `json:"documentSetVersions,omitempty"`
	DriveItem            *DriveItem            `json:"driveItem,omitempty"`
	ETag                 *string               `json:"eTag,omitempty"`
	Fields               *FieldValueSet        `json:"fields,omitempty"`
	Id                   *string               `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet          `json:"lastModifiedBy,omitempty"`
	LastModifiedByUser   *User                 `json:"lastModifiedByUser,omitempty"`
	LastModifiedDateTime *string               `json:"lastModifiedDateTime,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	ODataType            *string               `json:"@odata.type,omitempty"`
	ParentReference      *ItemReference        `json:"parentReference,omitempty"`
	SharepointIds        *SharepointIds        `json:"sharepointIds,omitempty"`
	Versions             *[]ListItemVersion    `json:"versions,omitempty"`
	WebUrl               *string               `json:"webUrl,omitempty"`
}
