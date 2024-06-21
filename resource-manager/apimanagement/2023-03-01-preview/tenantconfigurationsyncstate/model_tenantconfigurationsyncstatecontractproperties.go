package tenantconfigurationsyncstate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantConfigurationSyncStateContractProperties struct {
	Branch                  *string `json:"branch,omitempty"`
	CommitId                *string `json:"commitId,omitempty"`
	ConfigurationChangeDate *string `json:"configurationChangeDate,omitempty"`
	IsExport                *bool   `json:"isExport,omitempty"`
	IsGitEnabled            *bool   `json:"isGitEnabled,omitempty"`
	IsSynced                *bool   `json:"isSynced,omitempty"`
	LastOperationId         *string `json:"lastOperationId,omitempty"`
	SyncDate                *string `json:"syncDate,omitempty"`
}
