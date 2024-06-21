package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitLabProjectProperties struct {
	FullyQualifiedFriendlyName      *string                  `json:"fullyQualifiedFriendlyName,omitempty"`
	FullyQualifiedName              *string                  `json:"fullyQualifiedName,omitempty"`
	FullyQualifiedParentGroupName   *string                  `json:"fullyQualifiedParentGroupName,omitempty"`
	OnboardingState                 *OnboardingState         `json:"onboardingState,omitempty"`
	ProvisioningState               *DevOpsProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStatusMessage       *string                  `json:"provisioningStatusMessage,omitempty"`
	ProvisioningStatusUpdateTimeUtc *string                  `json:"provisioningStatusUpdateTimeUtc,omitempty"`
	Url                             *string                  `json:"url,omitempty"`
}
