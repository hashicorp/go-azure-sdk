package devops

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *AzureDevOpsProjectProperties) GetProvisioningStatusUpdateTimeUtcAsTime() (*time.Time, error) {
	if o.ProvisioningStatusUpdateTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ProvisioningStatusUpdateTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureDevOpsProjectProperties) SetProvisioningStatusUpdateTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ProvisioningStatusUpdateTimeUtc = &formatted
}
