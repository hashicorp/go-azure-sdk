package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AadUserNotificationRecipient struct {
	ODataType *string `json:"@odata.type,omitempty"`
	UserId    *string `json:"userId,omitempty"`
}
