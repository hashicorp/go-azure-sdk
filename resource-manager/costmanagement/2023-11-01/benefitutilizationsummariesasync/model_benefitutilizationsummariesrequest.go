package benefitutilizationsummariesasync

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummariesRequest struct {
	BenefitId        *string      `json:"benefitId,omitempty"`
	BenefitOrderId   *string      `json:"benefitOrderId,omitempty"`
	BillingAccountId *string      `json:"billingAccountId,omitempty"`
	BillingProfileId *string      `json:"billingProfileId,omitempty"`
	EndDate          string       `json:"endDate"`
	Grain            Grain        `json:"grain"`
	Kind             *BenefitKind `json:"kind,omitempty"`
	StartDate        string       `json:"startDate"`
}

func (o *BenefitUtilizationSummariesRequest) GetEndDateAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BenefitUtilizationSummariesRequest) SetEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDate = formatted
}

func (o *BenefitUtilizationSummariesRequest) GetStartDateAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BenefitUtilizationSummariesRequest) SetStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDate = formatted
}
