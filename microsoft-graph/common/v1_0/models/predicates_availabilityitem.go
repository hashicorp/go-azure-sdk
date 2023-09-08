package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilityItemOperationPredicate struct {
	ODataType *string
	ServiceId *string
}

func (p AvailabilityItemOperationPredicate) Matches(input AvailabilityItem) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ServiceId != nil && (input.ServiceId == nil || *p.ServiceId != *input.ServiceId) {
		return false
	}

	return true
}
