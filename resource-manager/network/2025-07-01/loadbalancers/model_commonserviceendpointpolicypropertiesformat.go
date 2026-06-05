package loadbalancers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonServiceEndpointPolicyPropertiesFormat struct {
	ContextualServiceEndpointPolicies *[]string                                `json:"contextualServiceEndpointPolicies,omitempty"`
	ProvisioningState                 *ProvisioningState                       `json:"provisioningState,omitempty"`
	ResourceGuid                      *string                                  `json:"resourceGuid,omitempty"`
	ServiceAlias                      *string                                  `json:"serviceAlias,omitempty"`
	ServiceEndpointPolicyDefinitions  *[]CommonServiceEndpointPolicyDefinition `json:"serviceEndpointPolicyDefinitions,omitempty"`
	Subnets                           *[]CommonSubnet                          `json:"subnets,omitempty"`
}
