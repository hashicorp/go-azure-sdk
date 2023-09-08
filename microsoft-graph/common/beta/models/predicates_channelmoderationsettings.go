package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelModerationSettingsOperationPredicate struct {
	AllowNewMessageFromBots       *bool
	AllowNewMessageFromConnectors *bool
	ODataType                     *string
}

func (p ChannelModerationSettingsOperationPredicate) Matches(input ChannelModerationSettings) bool {

	if p.AllowNewMessageFromBots != nil && (input.AllowNewMessageFromBots == nil || *p.AllowNewMessageFromBots != *input.AllowNewMessageFromBots) {
		return false
	}

	if p.AllowNewMessageFromConnectors != nil && (input.AllowNewMessageFromConnectors == nil || *p.AllowNewMessageFromConnectors != *input.AllowNewMessageFromConnectors) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
