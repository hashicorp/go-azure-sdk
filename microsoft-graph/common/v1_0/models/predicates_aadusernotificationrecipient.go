package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AadUserNotificationRecipientOperationPredicate struct {
	ODataType *string
	UserId    *string
}

func (p AadUserNotificationRecipientOperationPredicate) Matches(input AadUserNotificationRecipient) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
