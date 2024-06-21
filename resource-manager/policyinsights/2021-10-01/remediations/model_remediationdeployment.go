package remediations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemediationDeployment struct {
	CreatedOn            *string          `json:"createdOn,omitempty"`
	DeploymentId         *string          `json:"deploymentId,omitempty"`
	Error                *ErrorDefinition `json:"error,omitempty"`
	LastUpdatedOn        *string          `json:"lastUpdatedOn,omitempty"`
	RemediatedResourceId *string          `json:"remediatedResourceId,omitempty"`
	ResourceLocation     *string          `json:"resourceLocation,omitempty"`
	Status               *string          `json:"status,omitempty"`
}
