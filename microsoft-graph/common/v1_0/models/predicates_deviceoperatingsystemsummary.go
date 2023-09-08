package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceOperatingSystemSummaryOperationPredicate struct {
	AndroidCorporateWorkProfileCount *int64
	AndroidCount                     *int64
	AndroidDedicatedCount            *int64
	AndroidDeviceAdminCount          *int64
	AndroidFullyManagedCount         *int64
	AndroidWorkProfileCount          *int64
	ConfigMgrDeviceCount             *int64
	IosCount                         *int64
	MacOSCount                       *int64
	ODataType                        *string
	UnknownCount                     *int64
	WindowsCount                     *int64
	WindowsMobileCount               *int64
}

func (p DeviceOperatingSystemSummaryOperationPredicate) Matches(input DeviceOperatingSystemSummary) bool {

	if p.AndroidCorporateWorkProfileCount != nil && (input.AndroidCorporateWorkProfileCount == nil || *p.AndroidCorporateWorkProfileCount != *input.AndroidCorporateWorkProfileCount) {
		return false
	}

	if p.AndroidCount != nil && (input.AndroidCount == nil || *p.AndroidCount != *input.AndroidCount) {
		return false
	}

	if p.AndroidDedicatedCount != nil && (input.AndroidDedicatedCount == nil || *p.AndroidDedicatedCount != *input.AndroidDedicatedCount) {
		return false
	}

	if p.AndroidDeviceAdminCount != nil && (input.AndroidDeviceAdminCount == nil || *p.AndroidDeviceAdminCount != *input.AndroidDeviceAdminCount) {
		return false
	}

	if p.AndroidFullyManagedCount != nil && (input.AndroidFullyManagedCount == nil || *p.AndroidFullyManagedCount != *input.AndroidFullyManagedCount) {
		return false
	}

	if p.AndroidWorkProfileCount != nil && (input.AndroidWorkProfileCount == nil || *p.AndroidWorkProfileCount != *input.AndroidWorkProfileCount) {
		return false
	}

	if p.ConfigMgrDeviceCount != nil && (input.ConfigMgrDeviceCount == nil || *p.ConfigMgrDeviceCount != *input.ConfigMgrDeviceCount) {
		return false
	}

	if p.IosCount != nil && (input.IosCount == nil || *p.IosCount != *input.IosCount) {
		return false
	}

	if p.MacOSCount != nil && (input.MacOSCount == nil || *p.MacOSCount != *input.MacOSCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UnknownCount != nil && (input.UnknownCount == nil || *p.UnknownCount != *input.UnknownCount) {
		return false
	}

	if p.WindowsCount != nil && (input.WindowsCount == nil || *p.WindowsCount != *input.WindowsCount) {
		return false
	}

	if p.WindowsMobileCount != nil && (input.WindowsMobileCount == nil || *p.WindowsMobileCount != *input.WindowsMobileCount) {
		return false
	}

	return true
}
