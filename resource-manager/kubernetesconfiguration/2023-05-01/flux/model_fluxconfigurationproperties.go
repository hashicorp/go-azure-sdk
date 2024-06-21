package flux

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FluxConfigurationProperties struct {
	AzureBlob                      *AzureBlobDefinition                `json:"azureBlob,omitempty"`
	Bucket                         *BucketDefinition                   `json:"bucket,omitempty"`
	ComplianceState                *FluxComplianceState                `json:"complianceState,omitempty"`
	ConfigurationProtectedSettings *map[string]string                  `json:"configurationProtectedSettings,omitempty"`
	ErrorMessage                   *string                             `json:"errorMessage,omitempty"`
	GitRepository                  *GitRepositoryDefinition            `json:"gitRepository,omitempty"`
	Kustomizations                 *map[string]KustomizationDefinition `json:"kustomizations,omitempty"`
	Namespace                      *string                             `json:"namespace,omitempty"`
	ProvisioningState              *ProvisioningState                  `json:"provisioningState,omitempty"`
	ReconciliationWaitDuration     *string                             `json:"reconciliationWaitDuration,omitempty"`
	RepositoryPublicKey            *string                             `json:"repositoryPublicKey,omitempty"`
	Scope                          *ScopeType                          `json:"scope,omitempty"`
	SourceKind                     *SourceKindType                     `json:"sourceKind,omitempty"`
	SourceSyncedCommitId           *string                             `json:"sourceSyncedCommitId,omitempty"`
	SourceUpdatedAt                *string                             `json:"sourceUpdatedAt,omitempty"`
	StatusUpdatedAt                *string                             `json:"statusUpdatedAt,omitempty"`
	Statuses                       *[]ObjectStatusDefinition           `json:"statuses,omitempty"`
	Suspend                        *bool                               `json:"suspend,omitempty"`
	WaitForReconciliation          *bool                               `json:"waitForReconciliation,omitempty"`
}
