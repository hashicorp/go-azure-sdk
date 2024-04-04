package devops

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *GitHubRepositoryProperties) GetProvisioningStatusUpdateTimeUtcAsTime() (*time.Time, error) {
	if o.ProvisioningStatusUpdateTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ProvisioningStatusUpdateTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *GitHubRepositoryProperties) SetProvisioningStatusUpdateTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ProvisioningStatusUpdateTimeUtc = &formatted
}
