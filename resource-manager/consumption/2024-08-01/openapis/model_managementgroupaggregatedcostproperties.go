package openapis

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupAggregatedCostProperties struct {
	AzureCharges            *float64                               `json:"azureCharges,omitempty"`
	BillingPeriodId         *string                                `json:"billingPeriodId,omitempty"`
	ChargesBilledSeparately *float64                               `json:"chargesBilledSeparately,omitempty"`
	Children                *[]ManagementGroupAggregatedCostResult `json:"children,omitempty"`
	Currency                *string                                `json:"currency,omitempty"`
	ExcludedSubscriptions   *[]string                              `json:"excludedSubscriptions,omitempty"`
	IncludedSubscriptions   *[]string                              `json:"includedSubscriptions,omitempty"`
	MarketplaceCharges      *float64                               `json:"marketplaceCharges,omitempty"`
	UsageEnd                *string                                `json:"usageEnd,omitempty"`
	UsageStart              *string                                `json:"usageStart,omitempty"`
}

func (o *ManagementGroupAggregatedCostProperties) GetUsageEndAsTime() (*time.Time, error) {
	if o.UsageEnd == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UsageEnd, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagementGroupAggregatedCostProperties) SetUsageEndAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.UsageEnd = &formatted
}

func (o *ManagementGroupAggregatedCostProperties) GetUsageStartAsTime() (*time.Time, error) {
	if o.UsageStart == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UsageStart, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagementGroupAggregatedCostProperties) SetUsageStartAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.UsageStart = &formatted
}
