package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelUnsetAsFavoriteByDefaultEventMessageDetailOperationPredicate struct {
	ChannelId *string
	ODataType *string
}

func (p ChannelUnsetAsFavoriteByDefaultEventMessageDetailOperationPredicate) Matches(input ChannelUnsetAsFavoriteByDefaultEventMessageDetail) bool {

	if p.ChannelId != nil && (input.ChannelId == nil || *p.ChannelId != *input.ChannelId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
