package workspace

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceRepositoryConfiguration struct {
	AccountName         *string             `json:"accountName,omitempty"`
	ClientId            *string             `json:"clientId,omitempty"`
	ClientSecret        *GitHubClientSecret `json:"clientSecret,omitempty"`
	CollaborationBranch *string             `json:"collaborationBranch,omitempty"`
	HostName            *string             `json:"hostName,omitempty"`
	LastCommitId        *string             `json:"lastCommitId,omitempty"`
	ProjectName         *string             `json:"projectName,omitempty"`
	RepositoryName      *string             `json:"repositoryName,omitempty"`
	RootFolder          *string             `json:"rootFolder,omitempty"`
	TenantId            *string             `json:"tenantId,omitempty"`
	Type                *string             `json:"type,omitempty"`
}
