package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleOsVersionOperationPredicate struct {
	MaxOSVersion *string
	MinOSVersion *string
	Name         *string
	ODataType    *string
}

func (p DeviceManagementApplicabilityRuleOsVersionOperationPredicate) Matches(input DeviceManagementApplicabilityRuleOsVersion) bool {

	if p.MaxOSVersion != nil && (input.MaxOSVersion == nil || *p.MaxOSVersion != *input.MaxOSVersion) {
		return false
	}

	if p.MinOSVersion != nil && (input.MinOSVersion == nil || *p.MinOSVersion != *input.MinOSVersion) {
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
