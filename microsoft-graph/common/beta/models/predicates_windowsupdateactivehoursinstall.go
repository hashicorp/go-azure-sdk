package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateActiveHoursInstallOperationPredicate struct {
	ActiveHoursEnd   *string
	ActiveHoursStart *string
	ODataType        *string
}

func (p WindowsUpdateActiveHoursInstallOperationPredicate) Matches(input WindowsUpdateActiveHoursInstall) bool {

	if p.ActiveHoursEnd != nil && (input.ActiveHoursEnd == nil || *p.ActiveHoursEnd != *input.ActiveHoursEnd) {
		return false
	}

	if p.ActiveHoursStart != nil && (input.ActiveHoursStart == nil || *p.ActiveHoursStart != *input.ActiveHoursStart) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
