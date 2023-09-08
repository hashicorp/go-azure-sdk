package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActivity struct {
	Access           *AccessAction `json:"access,omitempty"`
	ActivityDateTime *string       `json:"activityDateTime,omitempty"`
	Actor            *IdentitySet  `json:"actor,omitempty"`
	DriveItem        *DriveItem    `json:"driveItem,omitempty"`
	Id               *string       `json:"id,omitempty"`
	ODataType        *string       `json:"@odata.type,omitempty"`
}
