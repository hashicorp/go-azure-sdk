package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterProperties struct {
	AadConfiguration                  *AadConfiguration                   `json:"aadConfiguration,omitempty"`
	AdministratorConfiguration        *AdministratorConfiguration         `json:"administratorConfiguration,omitempty"`
	AttachedNetworkIds                *[]string                           `json:"attachedNetworkIds,omitempty"`
	AvailableUpgrades                 *[]AvailableUpgrade                 `json:"availableUpgrades,omitempty"`
	ClusterId                         *string                             `json:"clusterId,omitempty"`
	ConnectedClusterId                *string                             `json:"connectedClusterId,omitempty"`
	ControlPlaneKubernetesVersion     *string                             `json:"controlPlaneKubernetesVersion,omitempty"`
	ControlPlaneNodeConfiguration     ControlPlaneNodeConfiguration       `json:"controlPlaneNodeConfiguration"`
	DetailedStatus                    *KubernetesClusterDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage             *string                             `json:"detailedStatusMessage,omitempty"`
	FeatureStatuses                   *[]FeatureStatus                    `json:"featureStatuses,omitempty"`
	InitialAgentPoolConfigurations    []InitialAgentPoolConfiguration     `json:"initialAgentPoolConfigurations"`
	KubernetesVersion                 string                              `json:"kubernetesVersion"`
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration  `json:"managedResourceGroupConfiguration,omitempty"`
	NetworkConfiguration              NetworkConfiguration                `json:"networkConfiguration"`
	Nodes                             *[]KubernetesClusterNode            `json:"nodes,omitempty"`
	ProvisioningState                 *KubernetesClusterProvisioningState `json:"provisioningState,omitempty"`
}
