package environments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentProperties struct {
	ArmTemplateDisplayName *string                          `json:"armTemplateDisplayName,omitempty"`
	CreatedByUser          *string                          `json:"createdByUser,omitempty"`
	DeploymentProperties   *EnvironmentDeploymentProperties `json:"deploymentProperties,omitempty"`
	ProvisioningState      *string                          `json:"provisioningState,omitempty"`
	ResourceGroupId        *string                          `json:"resourceGroupId,omitempty"`
	UniqueIdentifier       *string                          `json:"uniqueIdentifier,omitempty"`
}
