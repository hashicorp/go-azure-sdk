package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitHubOwnerProperties struct {
	GitHubInternalId                *string                  `json:"gitHubInternalId,omitempty"`
	OnboardingState                 *OnboardingState         `json:"onboardingState,omitempty"`
	OwnerUrl                        *string                  `json:"ownerUrl,omitempty"`
	ProvisioningState               *DevOpsProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStatusMessage       *string                  `json:"provisioningStatusMessage,omitempty"`
	ProvisioningStatusUpdateTimeUtc *string                  `json:"provisioningStatusUpdateTimeUtc,omitempty"`
}
