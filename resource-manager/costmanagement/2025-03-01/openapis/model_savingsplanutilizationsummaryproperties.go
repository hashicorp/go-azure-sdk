package openapis

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanUtilizationSummaryProperties struct {
	ArmSkuName               *string      `json:"armSkuName,omitempty"`
	AvgUtilizationPercentage *float64     `json:"avgUtilizationPercentage,omitempty"`
	BenefitId                *string      `json:"benefitId,omitempty"`
	BenefitOrderId           *string      `json:"benefitOrderId,omitempty"`
	BenefitType              *BenefitKind `json:"benefitType,omitempty"`
	MaxUtilizationPercentage *float64     `json:"maxUtilizationPercentage,omitempty"`
	MinUtilizationPercentage *float64     `json:"minUtilizationPercentage,omitempty"`
	UsageDate                *string      `json:"usageDate,omitempty"`
}

func (o *SavingsPlanUtilizationSummaryProperties) GetUsageDateAsTime() (*time.Time, error) {
	if o.UsageDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UsageDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SavingsPlanUtilizationSummaryProperties) SetUsageDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.UsageDate = &formatted
}
