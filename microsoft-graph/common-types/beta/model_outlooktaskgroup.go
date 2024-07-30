package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskGroup struct {
	ChangeKey      *string              `json:"changeKey,omitempty"`
	GroupKey       *string              `json:"groupKey,omitempty"`
	Id             *string              `json:"id,omitempty"`
	IsDefaultGroup *bool                `json:"isDefaultGroup,omitempty"`
	Name           *string              `json:"name,omitempty"`
	ODataType      *string              `json:"@odata.type,omitempty"`
	TaskFolders    *[]OutlookTaskFolder `json:"taskFolders,omitempty"`
}
