package managedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEnvironmentProperties struct {
	AppLogsConfiguration     *AppLogsConfiguration         `json:"appLogsConfiguration,omitempty"`
	DaprAIConnectionString   *string                       `json:"daprAIConnectionString,omitempty"`
	DaprAIInstrumentationKey *string                       `json:"daprAIInstrumentationKey,omitempty"`
	DefaultDomain            *string                       `json:"defaultDomain,omitempty"`
	DeploymentErrors         *string                       `json:"deploymentErrors,omitempty"`
	ProvisioningState        *EnvironmentProvisioningState `json:"provisioningState,omitempty"`
	StaticIP                 *string                       `json:"staticIp,omitempty"`
	VnetConfiguration        *VnetConfiguration            `json:"vnetConfiguration,omitempty"`
	ZoneRedundant            *bool                         `json:"zoneRedundant,omitempty"`
}
