package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceObjectiveCapability struct {
	ComputeModel                          *string                                  `json:"computeModel,omitempty"`
	Id                                    *string                                  `json:"id,omitempty"`
	IncludedMaxSize                       *MaxSizeCapability                       `json:"includedMaxSize,omitempty"`
	Name                                  *string                                  `json:"name,omitempty"`
	PerformanceLevel                      *PerformanceLevelCapability              `json:"performanceLevel,omitempty"`
	Reason                                *string                                  `json:"reason,omitempty"`
	Sku                                   *Sku                                     `json:"sku,omitempty"`
	Status                                *CapabilityStatus                        `json:"status,omitempty"`
	SupportedAutoPauseDelay               *AutoPauseDelayTimeRange                 `json:"supportedAutoPauseDelay,omitempty"`
	SupportedFreeLimitExhaustionBehaviors *[]FreeLimitExhaustionBehaviorCapability `json:"supportedFreeLimitExhaustionBehaviors,omitempty"`
	SupportedLicenseTypes                 *[]LicenseTypeCapability                 `json:"supportedLicenseTypes,omitempty"`
	SupportedMaintenanceConfigurations    *[]MaintenanceConfigurationCapability    `json:"supportedMaintenanceConfigurations,omitempty"`
	SupportedMaxSizes                     *[]MaxSizeRangeCapability                `json:"supportedMaxSizes,omitempty"`
	SupportedMinCapacities                *[]MinCapacityCapability                 `json:"supportedMinCapacities,omitempty"`
	ZonePinning                           *bool                                    `json:"zonePinning,omitempty"`
	ZoneRedundant                         *bool                                    `json:"zoneRedundant,omitempty"`
}
