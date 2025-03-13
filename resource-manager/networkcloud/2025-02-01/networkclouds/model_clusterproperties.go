package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProperties struct {
	AggregatorOrSingleRackDefinition  RackDefinition                     `json:"aggregatorOrSingleRackDefinition"`
	AnalyticsOutputSettings           *AnalyticsOutputSettings           `json:"analyticsOutputSettings,omitempty"`
	AnalyticsWorkspaceId              *string                            `json:"analyticsWorkspaceId,omitempty"`
	AvailableUpgradeVersions          *[]ClusterAvailableUpgradeVersion  `json:"availableUpgradeVersions,omitempty"`
	ClusterCapacity                   *ClusterCapacity                   `json:"clusterCapacity,omitempty"`
	ClusterConnectionStatus           *ClusterConnectionStatus           `json:"clusterConnectionStatus,omitempty"`
	ClusterExtendedLocation           *ExtendedLocation                  `json:"clusterExtendedLocation,omitempty"`
	ClusterLocation                   *string                            `json:"clusterLocation,omitempty"`
	ClusterManagerConnectionStatus    *ClusterManagerConnectionStatus    `json:"clusterManagerConnectionStatus,omitempty"`
	ClusterManagerId                  *string                            `json:"clusterManagerId,omitempty"`
	ClusterServicePrincipal           *ServicePrincipalInformation       `json:"clusterServicePrincipal,omitempty"`
	ClusterType                       ClusterType                        `json:"clusterType"`
	ClusterVersion                    string                             `json:"clusterVersion"`
	CommandOutputSettings             *CommandOutputSettings             `json:"commandOutputSettings,omitempty"`
	ComputeDeploymentThreshold        *ValidationThreshold               `json:"computeDeploymentThreshold,omitempty"`
	ComputeRackDefinitions            *[]RackDefinition                  `json:"computeRackDefinitions,omitempty"`
	DetailedStatus                    *ClusterDetailedStatus             `json:"detailedStatus,omitempty"`
	DetailedStatusMessage             *string                            `json:"detailedStatusMessage,omitempty"`
	HybridAksExtendedLocation         *ExtendedLocation                  `json:"hybridAksExtendedLocation,omitempty"`
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration `json:"managedResourceGroupConfiguration,omitempty"`
	ManualActionCount                 *int64                             `json:"manualActionCount,omitempty"`
	NetworkFabricId                   string                             `json:"networkFabricId"`
	ProvisioningState                 *ClusterProvisioningState          `json:"provisioningState,omitempty"`
	RuntimeProtectionConfiguration    *RuntimeProtectionConfiguration    `json:"runtimeProtectionConfiguration,omitempty"`
	SecretArchive                     *ClusterSecretArchive              `json:"secretArchive,omitempty"`
	SecretArchiveSettings             *SecretArchiveSettings             `json:"secretArchiveSettings,omitempty"`
	SupportExpiryDate                 *string                            `json:"supportExpiryDate,omitempty"`
	UpdateStrategy                    *ClusterUpdateStrategy             `json:"updateStrategy,omitempty"`
	VulnerabilityScanningSettings     *VulnerabilityScanningSettings     `json:"vulnerabilityScanningSettings,omitempty"`
	WorkloadResourceIds               *[]string                          `json:"workloadResourceIds,omitempty"`
}
