package sitepermission

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveRecipient struct {
	Alias     *string `json:"alias,omitempty"`
	Email     *string `json:"email,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	ObjectId  *string `json:"objectId,omitempty"`
}
