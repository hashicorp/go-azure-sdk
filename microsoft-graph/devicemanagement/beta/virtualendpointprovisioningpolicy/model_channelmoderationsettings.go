package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelModerationSettings struct {
	AllowNewMessageFromBots       *bool                                               `json:"allowNewMessageFromBots,omitempty"`
	AllowNewMessageFromConnectors *bool                                               `json:"allowNewMessageFromConnectors,omitempty"`
	ODataType                     *string                                             `json:"@odata.type,omitempty"`
	ReplyRestriction              *ChannelModerationSettingsReplyRestriction          `json:"replyRestriction,omitempty"`
	UserNewMessageRestriction     *ChannelModerationSettingsUserNewMessageRestriction `json:"userNewMessageRestriction,omitempty"`
}
