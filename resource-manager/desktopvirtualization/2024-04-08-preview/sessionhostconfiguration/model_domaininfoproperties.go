package sessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainInfoProperties struct {
	ActiveDirectoryInfo      *ActiveDirectoryInfoProperties      `json:"activeDirectoryInfo,omitempty"`
	AzureActiveDirectoryInfo *AzureActiveDirectoryInfoProperties `json:"azureActiveDirectoryInfo,omitempty"`
	JoinType                 DomainJoinType                      `json:"joinType"`
}
