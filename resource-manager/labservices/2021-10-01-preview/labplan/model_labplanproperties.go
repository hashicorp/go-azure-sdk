package labplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabPlanProperties struct {
	AllowedRegions             *[]string              `json:"allowedRegions,omitempty"`
	DefaultAutoShutdownProfile *AutoShutdownProfile   `json:"defaultAutoShutdownProfile"`
	DefaultConnectionProfile   *ConnectionProfile     `json:"defaultConnectionProfile"`
	DefaultNetworkProfile      *LabPlanNetworkProfile `json:"defaultNetworkProfile"`
	LinkedLmsInstance          *string                `json:"linkedLmsInstance,omitempty"`
	ProvisioningState          *ProvisioningState     `json:"provisioningState,omitempty"`
	SharedGalleryId            *string                `json:"sharedGalleryId,omitempty"`
	SupportInfo                *SupportInfo           `json:"supportInfo"`
}
