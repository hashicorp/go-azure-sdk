package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterPatchProperties struct {
	AggregatorOrSingleRackDefinition *RackDefinition                     `json:"aggregatorOrSingleRackDefinition,omitempty"`
	AnalyticsOutputSettings          *AnalyticsOutputSettings            `json:"analyticsOutputSettings,omitempty"`
	ClusterLocation                  *string                             `json:"clusterLocation,omitempty"`
	ClusterServicePrincipal          *ServicePrincipalInformation        `json:"clusterServicePrincipal,omitempty"`
	CommandOutputSettings            *CommandOutputSettings              `json:"commandOutputSettings,omitempty"`
	ComputeDeploymentThreshold       *ValidationThreshold                `json:"computeDeploymentThreshold,omitempty"`
	ComputeRackDefinitions           *[]RackDefinition                   `json:"computeRackDefinitions,omitempty"`
	RuntimeProtectionConfiguration   *RuntimeProtectionConfiguration     `json:"runtimeProtectionConfiguration,omitempty"`
	SecretArchive                    *ClusterSecretArchive               `json:"secretArchive,omitempty"`
	SecretArchiveSettings            *SecretArchiveSettings              `json:"secretArchiveSettings,omitempty"`
	UpdateStrategy                   *ClusterUpdateStrategy              `json:"updateStrategy,omitempty"`
	VulnerabilityScanningSettings    *VulnerabilityScanningSettingsPatch `json:"vulnerabilityScanningSettings,omitempty"`
}
