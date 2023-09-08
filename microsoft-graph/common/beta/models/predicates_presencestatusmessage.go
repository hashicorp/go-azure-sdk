package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PresenceStatusMessageOperationPredicate struct {
	ODataType         *string
	PublishedDateTime *string
}

func (p PresenceStatusMessageOperationPredicate) Matches(input PresenceStatusMessage) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PublishedDateTime != nil && (input.PublishedDateTime == nil || *p.PublishedDateTime != *input.PublishedDateTime) {
		return false
	}

	return true
}
