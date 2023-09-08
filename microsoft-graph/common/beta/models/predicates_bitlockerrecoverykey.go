package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitlockerRecoveryKeyOperationPredicate struct {
	CreatedDateTime *string
	DeviceId        *string
	Id              *string
	Key             *string
	ODataType       *string
}

func (p BitlockerRecoveryKeyOperationPredicate) Matches(input BitlockerRecoveryKey) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Key != nil && (input.Key == nil || *p.Key != *input.Key) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
