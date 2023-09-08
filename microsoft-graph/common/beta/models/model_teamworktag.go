package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkTag struct {
	Description *string              `json:"description,omitempty"`
	DisplayName *string              `json:"displayName,omitempty"`
	Id          *string              `json:"id,omitempty"`
	MemberCount *int64               `json:"memberCount,omitempty"`
	Members     *[]TeamworkTagMember `json:"members,omitempty"`
	ODataType   *string              `json:"@odata.type,omitempty"`
	TagType     *TeamworkTagTagType  `json:"tagType,omitempty"`
	TeamId      *string              `json:"teamId,omitempty"`
}
