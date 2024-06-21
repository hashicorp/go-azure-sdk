package workflowversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowVersionProperties struct {
	AccessControl          *FlowAccessControlConfiguration `json:"accessControl,omitempty"`
	AccessEndpoint         *string                         `json:"accessEndpoint,omitempty"`
	ChangedTime            *string                         `json:"changedTime,omitempty"`
	CreatedTime            *string                         `json:"createdTime,omitempty"`
	Definition             *interface{}                    `json:"definition,omitempty"`
	EndpointsConfiguration *FlowEndpointsConfiguration     `json:"endpointsConfiguration,omitempty"`
	IntegrationAccount     *ResourceReference              `json:"integrationAccount,omitempty"`
	Parameters             *map[string]WorkflowParameter   `json:"parameters,omitempty"`
	ProvisioningState      *WorkflowProvisioningState      `json:"provisioningState,omitempty"`
	Sku                    *Sku                            `json:"sku,omitempty"`
	State                  *WorkflowState                  `json:"state,omitempty"`
	Version                *string                         `json:"version,omitempty"`
}
