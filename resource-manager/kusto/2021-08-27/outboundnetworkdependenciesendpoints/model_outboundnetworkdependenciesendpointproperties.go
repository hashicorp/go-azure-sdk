package outboundnetworkdependenciesendpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutboundNetworkDependenciesEndpointProperties struct {
	Category          *string               `json:"category,omitempty"`
	Endpoints         *[]EndpointDependency `json:"endpoints,omitempty"`
	ProvisioningState *ProvisioningState    `json:"provisioningState,omitempty"`
}
