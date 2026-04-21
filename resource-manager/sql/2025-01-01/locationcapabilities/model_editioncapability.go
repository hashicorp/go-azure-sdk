package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionCapability struct {
	Name                            *string                       `json:"name,omitempty"`
	ReadScale                       *ReadScaleCapability          `json:"readScale,omitempty"`
	Reason                          *string                       `json:"reason,omitempty"`
	Status                          *CapabilityStatus             `json:"status,omitempty"`
	SupportedServiceLevelObjectives *[]ServiceObjectiveCapability `json:"supportedServiceLevelObjectives,omitempty"`
	SupportedStorageCapabilities    *[]StorageCapability          `json:"supportedStorageCapabilities,omitempty"`
	ZonePinning                     *bool                         `json:"zonePinning,omitempty"`
	ZoneRedundant                   *bool                         `json:"zoneRedundant,omitempty"`
}
