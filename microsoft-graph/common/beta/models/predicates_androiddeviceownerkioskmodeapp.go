package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeAppOperationPredicate struct {
	ClassName *string
	ODataType *string
	Package   *string
}

func (p AndroidDeviceOwnerKioskModeAppOperationPredicate) Matches(input AndroidDeviceOwnerKioskModeApp) bool {

	if p.ClassName != nil && (input.ClassName == nil || *p.ClassName != *input.ClassName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Package != nil && (input.Package == nil || *p.Package != *input.Package) {
		return false
	}

	return true
}
