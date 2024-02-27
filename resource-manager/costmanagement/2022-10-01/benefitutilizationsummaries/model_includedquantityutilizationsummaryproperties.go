package benefitutilizationsummaries

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludedQuantityUtilizationSummaryProperties struct {
	ArmSkuName            *string      `json:"armSkuName,omitempty"`
	BenefitId             *string      `json:"benefitId,omitempty"`
	BenefitOrderId        *string      `json:"benefitOrderId,omitempty"`
	BenefitType           *BenefitKind `json:"benefitType,omitempty"`
	UsageDate             *string      `json:"usageDate,omitempty"`
	UtilizationPercentage *float64     `json:"utilizationPercentage,omitempty"`
}

func (o *IncludedQuantityUtilizationSummaryProperties) GetUsageDateAsTime() (*time.Time, error) {
	if o.UsageDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UsageDate, "2006-01-02T15:04:05Z07:00")
}
