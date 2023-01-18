package integrationserviceenvironmentnetworkhealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentSubnetNetworkHealth struct {
	NetworkDependencyHealthState IntegrationServiceEnvironmentNetworkEndPointAccessibilityState `json:"networkDependencyHealthState"`
	OutboundNetworkDependencies  *[]IntegrationServiceEnvironmentNetworkDependency              `json:"outboundNetworkDependencies,omitempty"`
	OutboundNetworkHealth        *IntegrationServiceEnvironmentNetworkDependencyHealth          `json:"outboundNetworkHealth,omitempty"`
}
