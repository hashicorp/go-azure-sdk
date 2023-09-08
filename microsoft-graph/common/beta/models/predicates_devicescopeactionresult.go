package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeActionResultOperationPredicate struct {
	DeviceScopeAction *string
	DeviceScopeId     *string
	FailedMessage     *string
	ODataType         *string
}

func (p DeviceScopeActionResultOperationPredicate) Matches(input DeviceScopeActionResult) bool {

	if p.DeviceScopeAction != nil && (input.DeviceScopeAction == nil || *p.DeviceScopeAction != *input.DeviceScopeAction) {
		return false
	}

	if p.DeviceScopeId != nil && (input.DeviceScopeId == nil || *p.DeviceScopeId != *input.DeviceScopeId) {
		return false
	}

	if p.FailedMessage != nil && (input.FailedMessage == nil || *p.FailedMessage != *input.FailedMessage) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
