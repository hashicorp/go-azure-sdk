package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsNotificationOperationPredicate struct {
	ODataType   *string
	ResourceUrl *string
}

func (p CommsNotificationOperationPredicate) Matches(input CommsNotification) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourceUrl != nil && (input.ResourceUrl == nil || *p.ResourceUrl != *input.ResourceUrl) {
		return false
	}

	return true
}
