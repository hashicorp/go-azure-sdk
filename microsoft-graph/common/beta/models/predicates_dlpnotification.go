package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpNotificationOperationPredicate struct {
	Author    *string
	ODataType *string
}

func (p DlpNotificationOperationPredicate) Matches(input DlpNotification) bool {

	if p.Author != nil && (input.Author == nil || *p.Author != *input.Author) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
