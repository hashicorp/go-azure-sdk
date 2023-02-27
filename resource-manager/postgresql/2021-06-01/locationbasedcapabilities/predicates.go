package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityPropertiesOperationPredicate struct {
	GeoBackupSupported                   *bool
	Status                               *string
	ZoneRedundantHaAndGeoBackupSupported *bool
	ZoneRedundantHaSupported             *bool
}

func (p CapabilityPropertiesOperationPredicate) Matches(input CapabilityProperties) bool {

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
