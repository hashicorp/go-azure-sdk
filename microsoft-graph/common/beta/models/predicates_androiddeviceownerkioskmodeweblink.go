package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeWeblinkOperationPredicate struct {
	Label     *string
	Link      *string
	ODataType *string
}

func (p AndroidDeviceOwnerKioskModeWeblinkOperationPredicate) Matches(input AndroidDeviceOwnerKioskModeWeblink) bool {

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.Link != nil && (input.Link == nil || *p.Link != *input.Link) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
