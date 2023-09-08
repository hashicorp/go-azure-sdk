package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelAddedEventMessageDetailOperationPredicate struct {
	ChannelDisplayName *string
	ChannelId          *string
	ODataType          *string
}

func (p ChannelAddedEventMessageDetailOperationPredicate) Matches(input ChannelAddedEventMessageDetail) bool {

	if p.ChannelDisplayName != nil && (input.ChannelDisplayName == nil || *p.ChannelDisplayName != *input.ChannelDisplayName) {
		return false
	}

	if p.ChannelId != nil && (input.ChannelId == nil || *p.ChannelId != *input.ChannelId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
