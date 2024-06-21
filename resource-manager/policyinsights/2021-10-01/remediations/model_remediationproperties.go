package remediations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemediationProperties struct {
	CorrelationId               *string                                `json:"correlationId,omitempty"`
	CreatedOn                   *string                                `json:"createdOn,omitempty"`
	DeploymentStatus            *RemediationDeploymentSummary          `json:"deploymentStatus,omitempty"`
	FailureThreshold            *RemediationPropertiesFailureThreshold `json:"failureThreshold,omitempty"`
	Filters                     *RemediationFilters                    `json:"filters,omitempty"`
	LastUpdatedOn               *string                                `json:"lastUpdatedOn,omitempty"`
	ParallelDeployments         *int64                                 `json:"parallelDeployments,omitempty"`
	PolicyAssignmentId          *string                                `json:"policyAssignmentId,omitempty"`
	PolicyDefinitionReferenceId *string                                `json:"policyDefinitionReferenceId,omitempty"`
	ProvisioningState           *string                                `json:"provisioningState,omitempty"`
	ResourceCount               *int64                                 `json:"resourceCount,omitempty"`
	ResourceDiscoveryMode       *ResourceDiscoveryMode                 `json:"resourceDiscoveryMode,omitempty"`
	StatusMessage               *string                                `json:"statusMessage,omitempty"`
}
