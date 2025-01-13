package quota

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PTUDeploymentUsage struct {
	CollectionQuotaUsage *int64  `json:"collectionQuotaUsage,omitempty"`
	DeploymentName       *string `json:"deploymentName,omitempty"`
	ResourceGroup        *string `json:"resourceGroup,omitempty"`
	Usage                *int64  `json:"usage,omitempty"`
	WorkspaceName        *string `json:"workspaceName,omitempty"`
}
