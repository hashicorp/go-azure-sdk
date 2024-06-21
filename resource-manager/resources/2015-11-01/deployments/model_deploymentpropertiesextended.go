package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentPropertiesExtended struct {
	CorrelationId     *string         `json:"correlationId,omitempty"`
	Dependencies      *[]Dependency   `json:"dependencies,omitempty"`
	Error             *ErrorResponse  `json:"error,omitempty"`
	Mode              *DeploymentMode `json:"mode,omitempty"`
	Outputs           *interface{}    `json:"outputs,omitempty"`
	Parameters        *interface{}    `json:"parameters,omitempty"`
	ParametersLink    *ParametersLink `json:"parametersLink,omitempty"`
	Providers         *[]Provider     `json:"providers,omitempty"`
	ProvisioningState *string         `json:"provisioningState,omitempty"`
	Template          *interface{}    `json:"template,omitempty"`
	TemplateLink      *TemplateLink   `json:"templateLink,omitempty"`
	Timestamp         *string         `json:"timestamp,omitempty"`
}
