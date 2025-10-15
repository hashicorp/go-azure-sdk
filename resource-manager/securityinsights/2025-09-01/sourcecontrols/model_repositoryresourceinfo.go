package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RepositoryResourceInfo struct {
	AzureDevOpsResourceInfo *AzureDevOpsResourceInfo `json:"azureDevOpsResourceInfo,omitempty"`
	GitHubResourceInfo      *GitHubResourceInfo      `json:"gitHubResourceInfo,omitempty"`
	Webhook                 *Webhook                 `json:"webhook,omitempty"`
}
