package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallOptionsOperationPredicate struct {
	HideBotAfterEscalation              *bool
	IsContentSharingNotificationEnabled *bool
	ODataType                           *string
}

func (p CallOptionsOperationPredicate) Matches(input CallOptions) bool {

	if p.HideBotAfterEscalation != nil && (input.HideBotAfterEscalation == nil || *p.HideBotAfterEscalation != *input.HideBotAfterEscalation) {
		return false
	}

	if p.IsContentSharingNotificationEnabled != nil && (input.IsContentSharingNotificationEnabled == nil || *p.IsContentSharingNotificationEnabled != *input.IsContentSharingNotificationEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
