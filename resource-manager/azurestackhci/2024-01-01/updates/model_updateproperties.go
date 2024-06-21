package updates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateProperties struct {
	AdditionalProperties  *string                `json:"additionalProperties,omitempty"`
	AvailabilityType      *AvailabilityType      `json:"availabilityType,omitempty"`
	ComponentVersions     *[]PackageVersionInfo  `json:"componentVersions,omitempty"`
	Description           *string                `json:"description,omitempty"`
	DisplayName           *string                `json:"displayName,omitempty"`
	HealthCheckDate       *string                `json:"healthCheckDate,omitempty"`
	HealthCheckResult     *[]PrecheckResult      `json:"healthCheckResult,omitempty"`
	HealthState           *HealthState           `json:"healthState,omitempty"`
	InstalledDate         *string                `json:"installedDate,omitempty"`
	PackagePath           *string                `json:"packagePath,omitempty"`
	PackageSizeInMb       *float64               `json:"packageSizeInMb,omitempty"`
	PackageType           *string                `json:"packageType,omitempty"`
	Prerequisites         *[]UpdatePrerequisite  `json:"prerequisites,omitempty"`
	ProvisioningState     *ProvisioningState     `json:"provisioningState,omitempty"`
	Publisher             *string                `json:"publisher,omitempty"`
	RebootRequired        *RebootRequirement     `json:"rebootRequired,omitempty"`
	ReleaseLink           *string                `json:"releaseLink,omitempty"`
	State                 *State                 `json:"state,omitempty"`
	UpdateStateProperties *UpdateStateProperties `json:"updateStateProperties,omitempty"`
	Version               *string                `json:"version,omitempty"`
}
