package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInstallStateSummaryOperationPredicate struct {
	FailedDeviceCount       *int64
	Id                      *string
	InstalledDeviceCount    *int64
	NotInstalledDeviceCount *int64
	ODataType               *string
	UserName                *string
}

func (p UserInstallStateSummaryOperationPredicate) Matches(input UserInstallStateSummary) bool {

	if p.FailedDeviceCount != nil && (input.FailedDeviceCount == nil || *p.FailedDeviceCount != *input.FailedDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InstalledDeviceCount != nil && (input.InstalledDeviceCount == nil || *p.InstalledDeviceCount != *input.InstalledDeviceCount) {
		return false
	}

	if p.NotInstalledDeviceCount != nil && (input.NotInstalledDeviceCount == nil || *p.NotInstalledDeviceCount != *input.NotInstalledDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	return true
}
