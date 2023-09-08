package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EBookInstallSummaryOperationPredicate struct {
	FailedDeviceCount       *int64
	FailedUserCount         *int64
	Id                      *string
	InstalledDeviceCount    *int64
	InstalledUserCount      *int64
	NotInstalledDeviceCount *int64
	NotInstalledUserCount   *int64
	ODataType               *string
}

func (p EBookInstallSummaryOperationPredicate) Matches(input EBookInstallSummary) bool {

	if p.FailedDeviceCount != nil && (input.FailedDeviceCount == nil || *p.FailedDeviceCount != *input.FailedDeviceCount) {
		return false
	}

	if p.FailedUserCount != nil && (input.FailedUserCount == nil || *p.FailedUserCount != *input.FailedUserCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InstalledDeviceCount != nil && (input.InstalledDeviceCount == nil || *p.InstalledDeviceCount != *input.InstalledDeviceCount) {
		return false
	}

	if p.InstalledUserCount != nil && (input.InstalledUserCount == nil || *p.InstalledUserCount != *input.InstalledUserCount) {
		return false
	}

	if p.NotInstalledDeviceCount != nil && (input.NotInstalledDeviceCount == nil || *p.NotInstalledDeviceCount != *input.NotInstalledDeviceCount) {
		return false
	}

	if p.NotInstalledUserCount != nil && (input.NotInstalledUserCount == nil || *p.NotInstalledUserCount != *input.NotInstalledUserCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
