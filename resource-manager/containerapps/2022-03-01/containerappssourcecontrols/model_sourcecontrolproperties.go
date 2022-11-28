package containerappssourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlProperties struct {
	Branch                    *string                      `json:"branch,omitempty"`
	GitHubActionConfiguration *GithubActionConfiguration   `json:"githubActionConfiguration"`
	OperationState            *SourceControlOperationState `json:"operationState,omitempty"`
	RepoUrl                   *string                      `json:"repoUrl,omitempty"`
}
