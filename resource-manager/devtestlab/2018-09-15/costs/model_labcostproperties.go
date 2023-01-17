package costs

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabCostProperties struct {
	CreatedDate       *string                      `json:"createdDate,omitempty"`
	CurrencyCode      *string                      `json:"currencyCode,omitempty"`
	EndDateTime       *string                      `json:"endDateTime,omitempty"`
	LabCostDetails    *[]LabCostDetailsProperties  `json:"labCostDetails,omitempty"`
	LabCostSummary    *LabCostSummaryProperties    `json:"labCostSummary,omitempty"`
	ProvisioningState *string                      `json:"provisioningState,omitempty"`
	ResourceCosts     *[]LabResourceCostProperties `json:"resourceCosts,omitempty"`
	StartDateTime     *string                      `json:"startDateTime,omitempty"`
	TargetCost        *TargetCostProperties        `json:"targetCost,omitempty"`
	UniqueIdentifier  *string                      `json:"uniqueIdentifier,omitempty"`
}

func (o *LabCostProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LabCostProperties) SetCreatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDate = &formatted
}

func (o *LabCostProperties) GetEndDateTimeAsTime() (*time.Time, error) {
	if o.EndDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *LabCostProperties) SetEndDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDateTime = &formatted
}

func (o *LabCostProperties) GetStartDateTimeAsTime() (*time.Time, error) {
	if o.StartDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *LabCostProperties) SetStartDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDateTime = &formatted
}
