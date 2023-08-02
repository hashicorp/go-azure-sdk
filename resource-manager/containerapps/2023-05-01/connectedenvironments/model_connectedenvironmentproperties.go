package connectedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedEnvironmentProperties struct {
	CustomDomainConfiguration *CustomDomainConfiguration             `json:"customDomainConfiguration,omitempty"`
	DaprAIConnectionString    *string                                `json:"daprAIConnectionString,omitempty"`
	DefaultDomain             *string                                `json:"defaultDomain,omitempty"`
	DeploymentErrors          *string                                `json:"deploymentErrors,omitempty"`
	ProvisioningState         *ConnectedEnvironmentProvisioningState `json:"provisioningState,omitempty"`
	StaticIP                  *string                                `json:"staticIp,omitempty"`
}
