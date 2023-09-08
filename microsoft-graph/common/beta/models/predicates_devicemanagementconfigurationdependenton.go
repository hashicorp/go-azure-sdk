package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationDependentOnOperationPredicate struct {
	DependentOn     *string
	ODataType       *string
	ParentSettingId *string
}

func (p DeviceManagementConfigurationDependentOnOperationPredicate) Matches(input DeviceManagementConfigurationDependentOn) bool {

	if p.DependentOn != nil && (input.DependentOn == nil || *p.DependentOn != *input.DependentOn) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentSettingId != nil && (input.ParentSettingId == nil || *p.ParentSettingId != *input.ParentSettingId) {
		return false
	}

	return true
}
