package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicabilityOperationPredicate struct {
	Description *string
	ODataType   *string
}

func (p DeviceManagementConfigurationExchangeOnlineSettingApplicabilityOperationPredicate) Matches(input DeviceManagementConfigurationExchangeOnlineSettingApplicability) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
