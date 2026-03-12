package kubeenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubeEnvironmentPatchResourceProperties struct {
	AksResourceID               *string                           `json:"aksResourceID,omitempty"`
	AppLogsConfiguration        *AppLogsConfiguration             `json:"appLogsConfiguration,omitempty"`
	ArcConfiguration            *ArcConfiguration                 `json:"arcConfiguration,omitempty"`
	ContainerAppsConfiguration  *ContainerAppsConfiguration       `json:"containerAppsConfiguration,omitempty"`
	DefaultDomain               *string                           `json:"defaultDomain,omitempty"`
	DeploymentErrors            *string                           `json:"deploymentErrors,omitempty"`
	InternalLoadBalancerEnabled *bool                             `json:"internalLoadBalancerEnabled,omitempty"`
	ProvisioningState           *KubeEnvironmentProvisioningState `json:"provisioningState,omitempty"`
	StaticIP                    *string                           `json:"staticIp,omitempty"`
}
