// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package locationbasedcapabilities

type CapabilityPropertiesOperationPredicate struct {
	FastProvisioningSupported            *bool
	GeoBackupSupported                   *bool
	Status                               *string
	ZoneRedundantHaAndGeoBackupSupported *bool
	ZoneRedundantHaSupported             *bool
}

func (p CapabilityPropertiesOperationPredicate) Matches(input CapabilityProperties) bool {

	if p.FastProvisioningSupported != nil && (input.FastProvisioningSupported == nil && *p.FastProvisioningSupported != *input.FastProvisioningSupported) {
		return false
	}

	if p.GeoBackupSupported != nil && (input.GeoBackupSupported == nil && *p.GeoBackupSupported != *input.GeoBackupSupported) {
		return false
	}

	if p.Status != nil && (input.Status == nil && *p.Status != *input.Status) {
		return false
	}

	if p.ZoneRedundantHaAndGeoBackupSupported != nil && (input.ZoneRedundantHaAndGeoBackupSupported == nil && *p.ZoneRedundantHaAndGeoBackupSupported != *input.ZoneRedundantHaAndGeoBackupSupported) {
		return false
	}

	if p.ZoneRedundantHaSupported != nil && (input.ZoneRedundantHaSupported == nil && *p.ZoneRedundantHaSupported != *input.ZoneRedundantHaSupported) {
		return false
	}

	return true
}
