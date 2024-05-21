package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentProperties struct {
	Mode           *DeploymentMode `json:"mode,omitempty"`
	Parameters     *interface{}    `json:"parameters,omitempty"`
	ParametersLink *ParametersLink `json:"parametersLink,omitempty"`
	Template       *interface{}    `json:"template,omitempty"`
	TemplateLink   *TemplateLink   `json:"templateLink,omitempty"`
}
