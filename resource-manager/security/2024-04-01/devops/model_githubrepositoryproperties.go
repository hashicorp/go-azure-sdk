package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitHubRepositoryProperties struct {
	OnboardingState                 *OnboardingState         `json:"onboardingState,omitempty"`
	ParentOwnerName                 *string                  `json:"parentOwnerName,omitempty"`
	ProvisioningState               *DevOpsProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStatusMessage       *string                  `json:"provisioningStatusMessage,omitempty"`
	ProvisioningStatusUpdateTimeUtc *string                  `json:"provisioningStatusUpdateTimeUtc,omitempty"`
	RepoFullName                    *string                  `json:"repoFullName,omitempty"`
	RepoId                          *string                  `json:"repoId,omitempty"`
	RepoName                        *string                  `json:"repoName,omitempty"`
	RepoUrl                         *string                  `json:"repoUrl,omitempty"`
}
