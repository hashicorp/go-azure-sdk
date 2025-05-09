package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentValidateResult struct {
	Error      *ErrorResponse                `json:"error,omitempty"`
	Id         *string                       `json:"id,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *DeploymentPropertiesExtended `json:"properties,omitempty"`
	Type       *string                       `json:"type,omitempty"`
}
