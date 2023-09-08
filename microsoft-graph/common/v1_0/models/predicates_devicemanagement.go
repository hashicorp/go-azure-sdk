package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementOperationPredicate struct {
	Id              *string
	IntuneAccountId *string
	ODataType       *string
}

func (p DeviceManagementOperationPredicate) Matches(input DeviceManagement) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IntuneAccountId != nil && (input.IntuneAccountId == nil || *p.IntuneAccountId != *input.IntuneAccountId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
