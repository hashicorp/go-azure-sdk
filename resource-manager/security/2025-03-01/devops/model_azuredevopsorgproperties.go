package devops

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDevOpsOrgProperties struct {
	ActionableRemediation           *ActionableRemediation   `json:"actionableRemediation,omitempty"`
	OnboardingState                 *OnboardingState         `json:"onboardingState,omitempty"`
	ProvisioningState               *DevOpsProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStatusMessage       *string                  `json:"provisioningStatusMessage,omitempty"`
	ProvisioningStatusUpdateTimeUtc *string                  `json:"provisioningStatusUpdateTimeUtc,omitempty"`
}

func (o *AzureDevOpsOrgProperties) GetProvisioningStatusUpdateTimeUtcAsTime() (*time.Time, error) {
	if o.ProvisioningStatusUpdateTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ProvisioningStatusUpdateTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureDevOpsOrgProperties) SetProvisioningStatusUpdateTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ProvisioningStatusUpdateTimeUtc = &formatted
}
