package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelAddedEventMessageDetail struct {
	ChannelDisplayName *string      `json:"channelDisplayName,omitempty"`
	ChannelId          *string      `json:"channelId,omitempty"`
	Initiator          *IdentitySet `json:"initiator,omitempty"`
	ODataType          *string      `json:"@odata.type,omitempty"`
}
