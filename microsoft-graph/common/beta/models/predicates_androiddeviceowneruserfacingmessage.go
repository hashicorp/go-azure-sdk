package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerUserFacingMessageOperationPredicate struct {
	DefaultMessage *string
	ODataType      *string
}

func (p AndroidDeviceOwnerUserFacingMessageOperationPredicate) Matches(input AndroidDeviceOwnerUserFacingMessage) bool {

	if p.DefaultMessage != nil && (input.DefaultMessage == nil || *p.DefaultMessage != *input.DefaultMessage) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
