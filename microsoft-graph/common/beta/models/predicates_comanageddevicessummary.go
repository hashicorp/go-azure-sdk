package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDevicesSummaryOperationPredicate struct {
	CompliancePolicyCount         *int64
	ConfigurationSettingsCount    *int64
	EndpointProtectionCount       *int64
	InventoryCount                *int64
	ModernAppsCount               *int64
	ODataType                     *string
	OfficeAppsCount               *int64
	ResourceAccessCount           *int64
	TotalComanagedCount           *int64
	WindowsUpdateForBusinessCount *int64
}

func (p ComanagedDevicesSummaryOperationPredicate) Matches(input ComanagedDevicesSummary) bool {

	if p.CompliancePolicyCount != nil && (input.CompliancePolicyCount == nil || *p.CompliancePolicyCount != *input.CompliancePolicyCount) {
		return false
	}

	if p.ConfigurationSettingsCount != nil && (input.ConfigurationSettingsCount == nil || *p.ConfigurationSettingsCount != *input.ConfigurationSettingsCount) {
		return false
	}

	if p.EndpointProtectionCount != nil && (input.EndpointProtectionCount == nil || *p.EndpointProtectionCount != *input.EndpointProtectionCount) {
		return false
	}

	if p.InventoryCount != nil && (input.InventoryCount == nil || *p.InventoryCount != *input.InventoryCount) {
		return false
	}

	if p.ModernAppsCount != nil && (input.ModernAppsCount == nil || *p.ModernAppsCount != *input.ModernAppsCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OfficeAppsCount != nil && (input.OfficeAppsCount == nil || *p.OfficeAppsCount != *input.OfficeAppsCount) {
		return false
	}

	if p.ResourceAccessCount != nil && (input.ResourceAccessCount == nil || *p.ResourceAccessCount != *input.ResourceAccessCount) {
		return false
	}

	if p.TotalComanagedCount != nil && (input.TotalComanagedCount == nil || *p.TotalComanagedCount != *input.TotalComanagedCount) {
		return false
	}

	if p.WindowsUpdateForBusinessCount != nil && (input.WindowsUpdateForBusinessCount == nil || *p.WindowsUpdateForBusinessCount != *input.WindowsUpdateForBusinessCount) {
		return false
	}

	return true
}
