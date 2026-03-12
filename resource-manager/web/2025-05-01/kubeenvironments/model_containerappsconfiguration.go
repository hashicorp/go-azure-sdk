package kubeenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsConfiguration struct {
	AppSubnetResourceId          *string `json:"appSubnetResourceId,omitempty"`
	ControlPlaneSubnetResourceId *string `json:"controlPlaneSubnetResourceId,omitempty"`
	DaprAIInstrumentationKey     *string `json:"daprAIInstrumentationKey,omitempty"`
	DockerBridgeCidr             *string `json:"dockerBridgeCidr,omitempty"`
	PlatformReservedCidr         *string `json:"platformReservedCidr,omitempty"`
	PlatformReservedDnsIP        *string `json:"platformReservedDnsIP,omitempty"`
}
