package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeAppPositionItemOperationPredicate struct {
	ODataType *string
	Position  *int64
}

func (p AndroidDeviceOwnerKioskModeAppPositionItemOperationPredicate) Matches(input AndroidDeviceOwnerKioskModeAppPositionItem) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Position != nil && (input.Position == nil || *p.Position != *input.Position) {
		return false
	}

	return true
}
