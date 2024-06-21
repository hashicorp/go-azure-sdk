package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDevOpsRepositoryProperties struct {
	ActionableRemediation           *ActionableRemediation   `json:"actionableRemediation,omitempty"`
	OnboardingState                 *OnboardingState         `json:"onboardingState,omitempty"`
	ParentOrgName                   *string                  `json:"parentOrgName,omitempty"`
	ParentProjectName               *string                  `json:"parentProjectName,omitempty"`
	ProvisioningState               *DevOpsProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStatusMessage       *string                  `json:"provisioningStatusMessage,omitempty"`
	ProvisioningStatusUpdateTimeUtc *string                  `json:"provisioningStatusUpdateTimeUtc,omitempty"`
	RepoId                          *string                  `json:"repoId,omitempty"`
	RepoUrl                         *string                  `json:"repoUrl,omitempty"`
	Visibility                      *string                  `json:"visibility,omitempty"`
}
