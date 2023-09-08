package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppAssignmentSettingsOperationPredicate struct {
	ODataType          *string
	UseDeviceLicensing *bool
	VpnConfigurationId *string
}

func (p IosVppAppAssignmentSettingsOperationPredicate) Matches(input IosVppAppAssignmentSettings) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UseDeviceLicensing != nil && (input.UseDeviceLicensing == nil || *p.UseDeviceLicensing != *input.UseDeviceLicensing) {
		return false
	}

	if p.VpnConfigurationId != nil && (input.VpnConfigurationId == nil || *p.VpnConfigurationId != *input.VpnConfigurationId) {
		return false
	}

	return true
}
