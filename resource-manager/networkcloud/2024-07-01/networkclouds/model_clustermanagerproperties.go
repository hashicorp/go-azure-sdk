package networkclouds

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterManagerProperties struct {
	AnalyticsWorkspaceId              *string                            `json:"analyticsWorkspaceId,omitempty"`
	AvailabilityZones                 *zones.Schema                      `json:"availabilityZones,omitempty"`
	ClusterVersions                   *[]ClusterAvailableVersion         `json:"clusterVersions,omitempty"`
	DetailedStatus                    *ClusterManagerDetailedStatus      `json:"detailedStatus,omitempty"`
	DetailedStatusMessage             *string                            `json:"detailedStatusMessage,omitempty"`
	FabricControllerId                string                             `json:"fabricControllerId"`
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration `json:"managedResourceGroupConfiguration,omitempty"`
	ManagerExtendedLocation           *ExtendedLocation                  `json:"managerExtendedLocation,omitempty"`
	ProvisioningState                 *ClusterManagerProvisioningState   `json:"provisioningState,omitempty"`
	VMSize                            *string                            `json:"vmSize,omitempty"`
}
