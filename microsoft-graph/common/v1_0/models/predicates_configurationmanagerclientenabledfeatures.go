package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientEnabledFeaturesOperationPredicate struct {
	CompliancePolicy         *bool
	DeviceConfiguration      *bool
	Inventory                *bool
	ModernApps               *bool
	ODataType                *string
	ResourceAccess           *bool
	WindowsUpdateForBusiness *bool
}

func (p ConfigurationManagerClientEnabledFeaturesOperationPredicate) Matches(input ConfigurationManagerClientEnabledFeatures) bool {

	if p.CompliancePolicy != nil && (input.CompliancePolicy == nil || *p.CompliancePolicy != *input.CompliancePolicy) {
		return false
	}

	if p.DeviceConfiguration != nil && (input.DeviceConfiguration == nil || *p.DeviceConfiguration != *input.DeviceConfiguration) {
		return false
	}

	if p.Inventory != nil && (input.Inventory == nil || *p.Inventory != *input.Inventory) {
		return false
	}

	if p.ModernApps != nil && (input.ModernApps == nil || *p.ModernApps != *input.ModernApps) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourceAccess != nil && (input.ResourceAccess == nil || *p.ResourceAccess != *input.ResourceAccess) {
		return false
	}

	if p.WindowsUpdateForBusiness != nil && (input.WindowsUpdateForBusiness == nil || *p.WindowsUpdateForBusiness != *input.WindowsUpdateForBusiness) {
		return false
	}

	return true
}
