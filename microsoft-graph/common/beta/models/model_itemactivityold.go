package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActivityOLD struct {
	Action    *ItemActionSet       `json:"action,omitempty"`
	Actor     *IdentitySet         `json:"actor,omitempty"`
	DriveItem *DriveItem           `json:"driveItem,omitempty"`
	Id        *string              `json:"id,omitempty"`
	ListItem  *ListItem            `json:"listItem,omitempty"`
	ODataType *string              `json:"@odata.type,omitempty"`
	Times     *ItemActivityTimeSet `json:"times,omitempty"`
}
