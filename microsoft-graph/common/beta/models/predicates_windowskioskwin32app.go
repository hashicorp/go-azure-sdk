package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskWin32AppOperationPredicate struct {
	AutoLaunch                  *bool
	ClassicAppPath              *string
	EdgeKiosk                   *string
	EdgeKioskIdleTimeoutMinutes *int64
	EdgeNoFirstRun              *bool
	Name                        *string
	ODataType                   *string
}

func (p WindowsKioskWin32AppOperationPredicate) Matches(input WindowsKioskWin32App) bool {

	if p.AutoLaunch != nil && (input.AutoLaunch == nil || *p.AutoLaunch != *input.AutoLaunch) {
		return false
	}

	if p.ClassicAppPath != nil && (input.ClassicAppPath == nil || *p.ClassicAppPath != *input.ClassicAppPath) {
		return false
	}

	if p.EdgeKiosk != nil && (input.EdgeKiosk == nil || *p.EdgeKiosk != *input.EdgeKiosk) {
		return false
	}

	if p.EdgeKioskIdleTimeoutMinutes != nil && (input.EdgeKioskIdleTimeoutMinutes == nil || *p.EdgeKioskIdleTimeoutMinutes != *input.EdgeKioskIdleTimeoutMinutes) {
		return false
	}

	if p.EdgeNoFirstRun != nil && (input.EdgeNoFirstRun == nil || *p.EdgeNoFirstRun != *input.EdgeNoFirstRun) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
