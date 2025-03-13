package devops

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *GitLabProjectProperties) GetProvisioningStatusUpdateTimeUtcAsTime() (*time.Time, error) {
	if o.ProvisioningStatusUpdateTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ProvisioningStatusUpdateTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *GitLabProjectProperties) SetProvisioningStatusUpdateTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ProvisioningStatusUpdateTimeUtc = &formatted
}
