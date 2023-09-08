package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpWindowsDevicesNotificationOperationPredicate struct {
	Author        *string
	ContentName   *string
	LastModfiedBy *string
	ODataType     *string
}

func (p DlpWindowsDevicesNotificationOperationPredicate) Matches(input DlpWindowsDevicesNotification) bool {

	if p.Author != nil && (input.Author == nil || *p.Author != *input.Author) {
		return false
	}

	if p.ContentName != nil && (input.ContentName == nil || *p.ContentName != *input.ContentName) {
		return false
	}

	if p.LastModfiedBy != nil && (input.LastModfiedBy == nil || *p.LastModfiedBy != *input.LastModfiedBy) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
