package agentsconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentConnectorProperties struct {
	DataConnectorType  *string                     `json:"dataConnectorType,omitempty"`
	DataSource         *string                     `json:"dataSource,omitempty"`
	DeploymentError    *string                     `json:"deploymentError,omitempty"`
	Endpoint           *string                     `json:"endpoint,omitempty"`
	ExtendedProperties *map[string]interface{}     `json:"extendedProperties,omitempty"`
	Identity           *string                     `json:"identity,omitempty"`
	ProvisioningState  *ConnectorProvisioningState `json:"provisioningState,omitempty"`
}
