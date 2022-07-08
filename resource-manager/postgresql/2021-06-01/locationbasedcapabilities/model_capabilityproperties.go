package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityProperties struct {
	GeoBackupSupported                   *bool                              `json:"geoBackupSupported,omitempty"`
	Status                               *string                            `json:"status,omitempty"`
	SupportedFlexibleServerEditions      *[]FlexibleServerEditionCapability `json:"supportedFlexibleServerEditions,omitempty"`
	SupportedHAMode                      *[]string                          `json:"supportedHAMode,omitempty"`
	SupportedHyperscaleNodeEditions      *[]HyperscaleNodeEditionCapability `json:"supportedHyperscaleNodeEditions,omitempty"`
	Zone                                 *string                            `json:"zone,omitempty"`
	ZoneRedundantHaAndGeoBackupSupported *bool                              `json:"zoneRedundantHaAndGeoBackupSupported,omitempty"`
	ZoneRedundantHaSupported             *bool                              `json:"zoneRedundantHaSupported,omitempty"`
}
