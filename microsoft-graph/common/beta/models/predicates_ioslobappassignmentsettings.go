package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosLobAppAssignmentSettingsOperationPredicate struct {
	IsRemovable              *bool
	ODataType                *string
	PreventManagedAppBackup  *bool
	UninstallOnDeviceRemoval *bool
	VpnConfigurationId       *string
}

func (p IosLobAppAssignmentSettingsOperationPredicate) Matches(input IosLobAppAssignmentSettings) bool {

	if p.IsRemovable != nil && (input.IsRemovable == nil || *p.IsRemovable != *input.IsRemovable) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PreventManagedAppBackup != nil && (input.PreventManagedAppBackup == nil || *p.PreventManagedAppBackup != *input.PreventManagedAppBackup) {
		return false
	}

	if p.UninstallOnDeviceRemoval != nil && (input.UninstallOnDeviceRemoval == nil || *p.UninstallOnDeviceRemoval != *input.UninstallOnDeviceRemoval) {
		return false
	}

	if p.VpnConfigurationId != nil && (input.VpnConfigurationId == nil || *p.VpnConfigurationId != *input.VpnConfigurationId) {
		return false
	}

	return true
}
