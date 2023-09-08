package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelMembersNotificationRecipient struct {
	ChannelId *string `json:"channelId,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	TeamId    *string `json:"teamId,omitempty"`
}
