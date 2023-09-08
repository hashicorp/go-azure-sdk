package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelDescriptionUpdatedEventMessageDetailOperationPredicate struct {
	ChannelDescription *string
	ChannelId          *string
	ODataType          *string
}

func (p ChannelDescriptionUpdatedEventMessageDetailOperationPredicate) Matches(input ChannelDescriptionUpdatedEventMessageDetail) bool {

	if p.ChannelDescription != nil && (input.ChannelDescription == nil || *p.ChannelDescription != *input.ChannelDescription) {
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
