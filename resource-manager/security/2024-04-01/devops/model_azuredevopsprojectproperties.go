package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDevOpsProjectProperties struct {
	ActionableRemediation           *ActionableRemediation   `json:"actionableRemediation,omitempty"`
	OnboardingState                 *OnboardingState         `json:"onboardingState,omitempty"`
	ParentOrgName                   *string                  `json:"parentOrgName,omitempty"`
	ProjectId                       *string                  `json:"projectId,omitempty"`
	ProvisioningState               *DevOpsProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStatusMessage       *string                  `json:"provisioningStatusMessage,omitempty"`
	ProvisioningStatusUpdateTimeUtc *string                  `json:"provisioningStatusUpdateTimeUtc,omitempty"`
}
